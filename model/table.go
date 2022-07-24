package model

import "strings"

type Table struct {
	table   string
	asName  string
	columns []Column
}

func NewTable(table string, asName string, columns []Column) Table {
	return Table{table: table, asName: asName, columns: columns}
}
func (t *Table) GetTableName() string {
	return t.table
}

func (t *Table) GetAsName() string {
	return t.asName
}

func (t *Table) GetColumns() []Column {
	return t.columns
}

func (t *Table) AddColumn(column Column) {
	for _, col := range t.columns {
		if column.column == col.column {
			return
		}
	}
	column.table = t.table
	t.columns = append(t.columns, column)
}

func (t *Table) HasColumn(name string) bool {
	found := false
	for _, column := range t.columns {
		if strings.EqualFold(column.column, name) || strings.EqualFold(column.asName, name) {
			found = true
			break
		}
	}
	return found
}

func (t *Table) String() string {
	output := t.table + ":" + t.asName + "\n"
	for _, column := range t.columns {
		output += "\t" + column.String() + "\n"
	}
	return output
}
