package model

type Column struct {
	table  string
	column string
	asName string
}

func (c *Column) GetTableName() string {
	return c.table
}

func (c *Column) SetTableName(name string) {
	c.table = name
}

func (c *Column) String() string {
	output := c.table + ":" + c.column
	if c.asName != "" {
		output += ":" + c.asName
	}
	return output
}
