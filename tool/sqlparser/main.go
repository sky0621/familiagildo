package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v4"
)

func main() {
	result, err := pg_query.Parse("SELECT * FROM guild WHERE id = $1;")
	if err != nil {
		panic(err)
	}

	fmt.Println()

	if result != nil {
		for _, stmt := range result.GetStmts() {
			s := stmt.GetStmt()
			if s != nil {
				ss := s.GetSelectStmt()
				if ss != nil {
					froms := ss.FromClause
					for _, from := range froms {
						n := from.GetNode()
						nv, ok := n.(*pg_query.Node_RangeVar)
						if ok {
							fmt.Println(nv.RangeVar.Relname)
						}
					}
				}
			}
		}
	}
}
