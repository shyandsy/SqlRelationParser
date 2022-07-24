package model

import "strings"

type Relation struct {
	sourceTable  string
	sourceColumn string
	joinedTable  string
	joinedColumn string
}

func NewRelation(sourceTable string, sourceColumn string, joinedTable string, joinedColumn string) Relation {
	return Relation{sourceTable: sourceTable, sourceColumn: sourceColumn, joinedTable: joinedTable, joinedColumn: joinedColumn}
}

func (tr *Relation) GetSourceTable() string {
	return tr.sourceTable
}
func (tr *Relation) GetSourceColumn() string {
	return tr.sourceColumn
}
func (tr *Relation) GetJoinedTable() string {
	return tr.joinedTable
}
func (tr *Relation) GetJoinedColumn() string {
	return tr.joinedColumn
}

func (tr *Relation) SetSourceTable(name string) {
	tr.sourceTable = name
}

func (tr *Relation) SetSourceColumn(name string) {
	tr.sourceColumn = name
}

func (tr *Relation) SetJoinedTable(name string) {
	tr.joinedTable = name
}

func (tr *Relation) SetJoinedColumn(name string) {
	tr.joinedColumn = name
}

func (tr *Relation) Equals(relation Relation) bool {
	if strings.EqualFold(tr.sourceTable, relation.sourceTable) &&
		strings.EqualFold(tr.sourceColumn, relation.sourceColumn) &&
		strings.EqualFold(tr.joinedTable, relation.joinedTable) &&
		strings.EqualFold(tr.joinedColumn, relation.joinedColumn) {
		return true
	}
	return false
}

func (tr *Relation) String() string {
	return tr.sourceTable + ":" + tr.sourceColumn + " => " + tr.joinedTable + ":" + tr.joinedColumn
}
