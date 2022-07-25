package main

import (
	"fmt"

	SqlRelationParser "github.com/shyandsy/SqlRelationParser"
)

func main() {
	query := []string{
		"SELECT b.id as bid, b.title, b.type, a.last_name AS author, t.last_name AS translator FROM books b " +
			"LEFT JOIN authors a ON b.author_id = a.id " +
			"LEFT JOIN translators t ON b.translator_id = t.id " +
			"ORDER BY b.id;",
		"select person_id from address adr " +
			"inner join option_address_type opt_adt on opt_adt.option_id = ? and opt_adt.type_id = adr.type_id " +
			"and opt_adt.sequence_id = adr.sequence_id and opt_adt.sequence_id=1",
	}

	parser := SqlRelationParser.NewSqlRelationParser()
	result, err := parser.ParseRelationFromBatchSql(query)
	if err != nil {
		fmt.Println("failed to parse")
		return
	}
	fmt.Println("\n===============================")
	relations := result.GetRelations()
	fmt.Println("parse the relations success")
	for _, relation := range relations {
		fmt.Println(&relation)
	}
}
