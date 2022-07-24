package model

import (
	"fmt"
	"strings"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/opcode"
)

type ParserResult struct {
	tables    []Table
	columns   []Column
	relations []Relation
}

func (r *ParserResult) Enter(in ast.Node) (ast.Node, bool) {
	if name, ok := in.(*ast.ColumnName); ok {
		column := Column{column: name.Name.L, table: name.Table.L}
		r.columns = append(r.columns, column)
		return in, false
	}
	if name, ok := in.(*ast.TableSource); ok {
		if table, ok := name.Source.(*ast.TableName); ok {
			t := Table{table: table.Name.L, asName: name.AsName.L}
			r.tables = append(r.tables, t)
			return in, false
		}
	}
	if onCondition, ok := in.(*ast.OnCondition); ok {
		if exp, ok := onCondition.Expr.(*ast.BinaryOperationExpr); ok && exp != nil {
			r.getRelation(exp)
		}
	}
	return in, false
}

func (r *ParserResult) getRelation(exp *ast.BinaryOperationExpr) {
	if exp.Op == opcode.EQ {
		var leftColumn, rightColumn *Column
		if columnExp, ok := exp.L.(*ast.ColumnNameExpr); ok {
			if columnExp.Name != nil {
				leftColumn = &Column{column: columnExp.Name.Name.L, table: columnExp.Name.Table.L}
			}
		}
		if columnExp, ok := exp.R.(*ast.ColumnNameExpr); ok {
			if columnExp.Name != nil {
				rightColumn = &Column{column: columnExp.Name.Name.L, table: columnExp.Name.Table.L}
			}
		}
		if leftColumn != nil && rightColumn != nil {
			r.relations = append(r.relations, Relation{
				sourceTable:  leftColumn.table,
				sourceColumn: leftColumn.column,
				joinedTable:  rightColumn.table,
				joinedColumn: rightColumn.column,
			})
		}
	} else if exp.Op == opcode.LogicAnd || exp.Op == opcode.LogicOr {
		if expL, ok := exp.L.(*ast.BinaryOperationExpr); ok && exp != nil {
			r.getRelation(expL)
		}
		if expR, ok := exp.R.(*ast.BinaryOperationExpr); ok && exp != nil {
			r.getRelation(expR)
		}
	}
}

func (r *ParserResult) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func (r *ParserResult) GetTables() []Table {
	return r.tables
}

func (r *ParserResult) GetColumns() []Column {
	return r.columns
}

func (r *ParserResult) GetRelations() []Relation {
	return r.relations
}

func (r *ParserResult) Show() {
	line := []string{"Tables"}
	for _, table := range r.tables {
		line = append(line, table.String())
	}
	line = append(line, "Column:")
	for _, column := range r.columns {
		line = append(line, column.String())
	}
	line = append(line, "relation:")
	for _, relation := range r.relations {
		line = append(line, relation.String())
	}
	fmt.Println(strings.Join(line, "\n"))
}
