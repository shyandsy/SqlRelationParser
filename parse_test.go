package sql

import (
	"log"

	"github.com/shyandsy/SqlRelationParser/model"

	"testing"
)

func TestParseTable(t *testing.T) {
	testCase1(t)
	testCase2(t)
	testCase3(t)
	testCase4(t)
}

func testCase1(t *testing.T) {
	query := "SELECT b.id as bid, b.title, b.type, a.last_name AS author, t.last_name AS translator FROM books b " +
		"LEFT JOIN authors a ON b.author_id = a.id " +
		"LEFT JOIN translators t ON b.translator_id = t.id " +
		"ORDER BY b.id;"

	parser := NewSqlRelationParser()
	result, err := parser.ParseRelation(query)
	if err != nil {
		t.Error("failed to parse")
	}

	tables := result.GetTables()
	if len(tables) != 3 {
		t.Error("wrong number of tables")
	}

	relations := result.GetRelations()
	correctRelation := []model.Relation{
		model.NewRelation("books", "author_id", "authors", "id"),
		model.NewRelation("books", "translator_id", "translators", "id"),
	}
	if !verifyRelation(relations, correctRelation) {
		t.Error("some relation is not correct")
	}
}

func testCase2(t *testing.T) {
	query := "select person_id from address adr " +
		"inner join option_address_type opt_adt on opt_adt.option_id = ? and opt_adt.type_id = adr.type_id " +
		"and opt_adt.sequence_id = adr.sequence_id and opt_adt.sequence_id=1"

	parser := NewSqlRelationParser()
	result, err := parser.ParseRelation(query)
	if err != nil {
		t.Error("failed to parse")
	}

	tables := result.GetTables()
	if len(tables) != 2 {
		t.Error("wrong number of tables")
	}

	relations := result.GetRelations()
	correctRelation := []model.Relation{
		model.NewRelation("option_address_type", "type_id", "address", "type_id"),
		model.NewRelation("option_address_type", "sequence_id", "address", "sequence_id"),
	}
	if !verifyRelation(relations, correctRelation) {
		t.Error("some relation is not correct")
	}
}

func testCase3(t *testing.T) {
	query := []string{
		"SELECT b.id as bid, b.title, b.type, a.last_name AS author, t.last_name AS translator FROM books b " +
			"LEFT JOIN authors a ON b.author_id = a.id " +
			"LEFT JOIN translators t ON b.translator_id = t.id " +
			"ORDER BY b.id;",
		"select person_id from address adr " +
			"inner join option_address_type opt_adt on opt_adt.option_id = ? and opt_adt.type_id = adr.type_id " +
			"and opt_adt.sequence_id = adr.sequence_id and opt_adt.sequence_id=1",
	}

	parser := NewSqlRelationParser()
	result, err := parser.ParseRelationFromBatchSql(query)
	if err != nil {
		t.Error("failed to parse")
	}

	tables := result.GetTables()
	if len(tables) != 5 {
		t.Error("wrong number of tables")
	}

	relations := result.GetRelations()
	correctRelation := []model.Relation{
		model.NewRelation("books", "author_id", "authors", "id"),
		model.NewRelation("books", "translator_id", "translators", "id"),
		model.NewRelation("option_address_type", "type_id", "address", "type_id"),
		model.NewRelation("option_address_type", "sequence_id", "address", "sequence_id"),
	}
	if !verifyRelation(relations, correctRelation) {
		t.Error("some relation is not correct")
	}
}

func testCase4(t *testing.T) {
	query := `update (
			ctarticle as a left join (ctnclass as c left join ctclass as d on c.classid = d.classid) on a.nclassid = c.nclassid and a.classid = c.classid
		) 
		left join cttag as b on a.articleid = b.articleid set tag=d.class+' '+c.nclass, b.articleid=a.articleid, b.classid=a.classid, b.nclassid=a.nclassid 
		where a.classid=23 and a.nclassid=197;`
	parser := NewSqlRelationParser()
	result, err := parser.ParseRelation(query)
	if err != nil {
		t.Error("failed to parse", err)
	}
	log.Println(result)

	tables := result.GetTables()
	if len(tables) != 4 {
		t.Error("wrong number of tables")
	}

	relations := result.GetRelations()
	correctRelation := []model.Relation{
		model.NewRelation("ctnclass", "classid", "ctclass", "classid"),
		model.NewRelation("ctarticle", "nclassid", "ctnclass", "nclassid"),
		model.NewRelation("ctarticle", "classid", "ctnclass", "classid"),
		model.NewRelation("ctarticle", "articleid", "cttag", "articleid"),
	}
	if !verifyRelation(relations, correctRelation) {
		t.Error("some relation is not correct")
	}
}

func verifyRelation(results []model.Relation, correctRelations []model.Relation) bool {
	if len(results) != len(correctRelations) {
		log.Println("the total number of relation is not correct")
		return false
	}
	for _, result := range results {
		for i := 0; i < len(correctRelations); i++ {
			if result.Equals(correctRelations[i]) {
				correctRelations = append(correctRelations[:i], correctRelations[i+1:]...)
			}
		}
	}
	if len(correctRelations) > 0 {
		log.Println("some relation is not correct")
		return false
	}

	return true
}
