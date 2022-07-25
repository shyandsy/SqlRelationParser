# SqlRelationParser
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go](https://github.com/shyandsy/SqlRelationParser/actions/workflows/go.yml/badge.svg?event=push)](https://github.com/shyandsy/SqlRelationParser/actions/workflows/go.yml)

SqlRelationParser is a open source component to parse relations from sql statement  

## Usage
get relations for a single sql statement
```go
package main

import (
	"fmt"

	SqlRelationParser "github.com/shyandsy/SqlRelationParser"
)

func main() {
	query := "select person_id from address adr " +
		"inner join option_address_type opt_adt on opt_adt.option_id = ? and opt_adt.type_id = adr.type_id " +
		"and opt_adt.sequence_id = adr.sequence_id and opt_adt.sequence_id=1"

	parser := SqlRelationParser.NewSqlRelationParser()
	result, err := parser.ParseRelation(query)
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
```

result
```
parse the relations success
option_address_type:type_id => address:type_id
option_address_type:sequence_id => address:sequence_id
```

get relations for a batch of sql statements
```go
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
```

result
```
Tables:
books:b
	books:id
	books:title
	books:type
	books:author_id
	books:translator_id
authors:a
	authors:last_name
	authors:id
translators:t
	translators:last_name
	translators:id
address:adr
	address:type_id
	address:sequence_id
option_address_type:opt_adt
	option_address_type:option_id
	option_address_type:type_id
	option_address_type:sequence_id
Relations:
books:author_id => authors:id
books:translator_id => translators:id
option_address_type:type_id => address:type_id
option_address_type:sequence_id => address:sequence_id
```
## Installation

```shell
go get github.com/shyandsy/SqlRelationParser@latest
```

## Examples
example can be found in the ***example*** folder for this package.

## Support
Welcome for your Issue and PR