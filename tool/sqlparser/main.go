package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v4"
)

const selectSQL01 = `
SELECT p.p_name, SUM(o.quantity), SUM(p.price * o.quantity)
FROM order_desc AS o RIGHT JOIN product AS p
ON p.p_id = o.p_id
GROUP BY p.p_id, p.p_name ORDER BY SUM(p.price * o.quantity) DESC;
`

const insertSQL01 = `
INSERT INTO guild (name, status) VALUES ($1, 1)
RETURNING *;
`

const updateSQL01 = `
UPDATE guild SET status = 2 WHERE id = $1
RETURNING *;
`

func main() {
	result, err := pg_query.Parse(selectSQL01)
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
				is := s.GetInsertStmt()
				if is != nil {
					rel := is.GetRelation()
					if rel != nil {
						fmt.Println(rel.GetRelname())
					}
				}
				us := s.GetUpdateStmt()
				if us != nil {
					rel := us.GetRelation()
					if rel != nil {
						fmt.Println(rel.GetRelname())
					}
				}
			}
		}
	}
}
