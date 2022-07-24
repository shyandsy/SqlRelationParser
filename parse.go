package sql

import (
	"fmt"
	"log"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/shyandsy/SqlRelationParser/model"

	_ "github.com/pingcap/tidb/parser/test_driver"
)

type RelationParser interface {
	ParseRelation(sql string) (*model.Schema, error)
}

type relationParser struct{}

func NewSqlRelationParser() RelationParser {
	return &relationParser{}
}

func (r relationParser) ParseRelation(sql string) (*model.Schema, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}
	fmt.Println(stmtNodes)

	result := extractField(&stmtNodes[0])
	fmt.Println(result)
	result.Show()

	return mergeParserResult(result), nil

}

func extractField(rootNode *ast.StmtNode) *model.ParserResult {
	result := &model.ParserResult{}
	(*rootNode).Accept(result)
	return result
}

func mergeParserResult(result *model.ParserResult) *model.Schema {
	schema := model.Schema{}

	tables := result.GetTables()
	for _, table := range tables {
		schema.AddTable(table)
	}

	columns := result.GetColumns()
	for _, column := range columns {
		table := schema.GetTable(column.GetTableName())
		if table == nil {
			log.Println("ignore column :" + column.String())
			continue
		}
		column.SetTableName(table.GetTableName())
		table.AddColumn(column)
	}

	relations := result.GetRelations()
	for _, relation := range relations {
		sourceTable := schema.GetTable(relation.GetSourceTable())
		if sourceTable == nil {
			log.Println("source table doesnt exist, ignore relation :" + relation.String())
			continue
		}
		joinedTable := schema.GetTable(relation.GetJoinedTable())
		if joinedTable == nil {
			log.Println("joined table doesnt exist, ignore relation :" + relation.String())
			continue
		}
		if !sourceTable.HasColumn(relation.GetSourceColumn()) {
			log.Println("source column doesnt exist, ignore relation :" + relation.String())
			continue
		}
		if !joinedTable.HasColumn(relation.GetJoinedColumn()) {
			log.Println("joined column doesnt exist, ignore relation :" + relation.String())
			continue
		}
		relation.SetSourceTable(sourceTable.GetTableName())
		relation.SetJoinedTable(joinedTable.GetTableName())

		schema.AddRelation(relation)
	}

	return &schema
}
