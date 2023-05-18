package main

import (
	"fmt"
	"testing"

	"github.com/sky0621/familiagildo/tool/sqlparser"
)

func Test_collectTableNames(t *testing.T) {
	sqlParseResults := []*sqlparser.SQLParseResult{
		{TableNameWithCRUDSlice: []*sqlparser.TableNameWithCRUD{
			{TableName: "C"},
			{TableName: "B"},
			{TableName: "C"},
		}},
		{TableNameWithCRUDSlice: []*sqlparser.TableNameWithCRUD{
			{TableName: "D1"},
			{TableName: "B"},
			{TableName: "D"},
			{TableName: "A"},
		}},
	}
	actual := collectTableNames(sqlParseResults)
	for _, x := range actual {
		fmt.Println(x)
	}
}
