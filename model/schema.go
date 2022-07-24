package model

import (
	"strings"
)

type Schema struct {
	tables    []Table
	relations []Relation
}

func (s *Schema) AddTable(table Table) {
	s.tables = append(s.tables, table)
}

func (s *Schema) GetTable(name string) *Table {
	var result *Table
	for i := range s.tables {
		if strings.EqualFold(s.tables[i].table, name) || strings.EqualFold(s.tables[i].asName, name) {
			result = &s.tables[i]
			break
		}
	}
	return result
}

func (s *Schema) AddRelation(relation Relation) {
	s.relations = append(s.relations, relation)
}

func (s *Schema) GetTables() []Table {
	return s.tables
}

func (s *Schema) GetRelations() []Relation {
	return s.relations
}

func (s *Schema) String() string {
	output := "-----------------------\n"
	output += "Tables:\n"
	for _, table := range s.tables {
		output += table.String()
	}
	output += "Relations:\n"
	for _, relation := range s.relations {
		output += relation.String() + "\n"
	}
	output += "-----------------------"
	return output
}
