package main

import (
	"errors"
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v4"
)

const selectSQL01 = `
SELECT p.p_name, SUM(o.quantity), SUM(p.price * o.quantity)
FROM order_desc AS o RIGHT JOIN product AS p
ON p.p_id = o.p_id
GROUP BY p.p_id, p.p_name ORDER BY SUM(p.price * o.quantity) DESC;
`

const selectSQL02 = `
SELECT COUNT(bp.product_id) AS how_many_products,
COUNT(dev.account_id) AS how_many_developers,
COUNT(b.bug_id)/COUNT(dev.account_id) AS avg_bugs_per_developer,
COUNT(cust.account_id) AS how_many_customers
FROM Bugs b JOIN BugsProducts bp ON (b.bug_id = bp.bug_id)
JOIN Accounts dev ON (b.assigned_to = dev.account_id)
JOIN Accounts cust ON (b.reported_by = cust.account_id)
WHERE cust.email NOT LIKE '%@example.com'
GROUP BY bp.product_id;
`

const insertSQL01 = `
INSERT INTO guild (name, status) VALUES ($1, 1)
RETURNING *;
`

const updateSQL01 = `
UPDATE guild SET status = 2 WHERE id = $1
RETURNING *;
`

const deleteSQL01 = `
DELETE FROM guest_token WHERE id = $1;
`

func main() {
	execMain()
}

func execMain() {
	//result, err := pg_query.Parse(selectSQL01)
	//result, err := pg_query.Parse(insertSQL01)
	//result, err := pg_query.Parse(updateSQL01)
	result, err := pg_query.Parse(deleteSQL01)
	if err != nil {
		panic(err)
	}

	/*	resJSON, err := pg_query.ParseToJSON(selectSQL01)
		if err != nil {
			panic(err)
		}
		fmt.Println(resJSON)
	*/

	if result == nil {
		panic(errors.New("result == nil"))
	}

	for _, stmt := range result.GetStmts() {
		result, err := parseStmt(stmt.GetStmt())
		if err != nil {
			fmt.Println("parseStmt is nil")
			continue
		}
		if result == nil {
			continue
		}
		fmt.Println(result.crud)
		for _, tw := range result.tableWiths {
			fmt.Println(tw.tableName)
		}
	}
}

// sql-name
//   table-name
//     crud(s)

type SQLName string

type TableName string

type CRUD int8

const (
	Select CRUD = iota + 1
	Insert
	Update
	Delete
)

type TableWith struct {
	tableName TableName
}

type CRUDTableNames struct {
	crud       CRUD
	tableWiths []TableWith
}

func parseStmt(s *pg_query.Node) (*CRUDTableNames, error) {
	if s == nil {
		return nil, errors.New("node is nil")
	}

	var result *CRUDTableNames
	var err error

	result, err = parseSelectStmt(s.GetSelectStmt())
	if err != nil {
		return nil, err
	}
	if result != nil {
		return result, nil
	}

	result, err = parseInsertStmt(s.GetInsertStmt())
	if err != nil {
		return nil, err
	}
	if result != nil {
		return result, nil
	}

	result, err = parseUpdateStmt(s.GetUpdateStmt())
	if err != nil {
		return nil, err
	}
	if result != nil {
		return result, nil
	}

	result, err = parseDeleteStmt(s.GetDeleteStmt())
	if err != nil {
		return nil, err
	}
	if result != nil {
		return result, nil
	}

	return nil, nil
}

func parseSelectStmt(s *pg_query.SelectStmt) (*CRUDTableNames, error) {
	if s == nil {
		return nil, nil
	}

	fromArray := s.FromClause
	if fromArray == nil {
		return nil, nil
	}

	result := &CRUDTableNames{crud: Select}

	for _, from := range fromArray {
		n := from.GetNode()
		nv, ok := n.(*pg_query.Node_RangeVar)
		if ok {
			if nv != nil && nv.RangeVar != nil {
				result.tableWiths = append(result.tableWiths, TableWith{tableName: TableName(nv.RangeVar.Relname)})
			}
		}
		nj, ok2 := n.(*pg_query.Node_JoinExpr)
		if ok2 {
			if nj != nil && nj.JoinExpr != nil {
				if nj.JoinExpr.Larg != nil && nj.JoinExpr.Larg.GetNode() != nil {
					nl := nj.JoinExpr.Larg.GetNode()
					nv, ok := nl.(*pg_query.Node_RangeVar)
					if ok {
						if nv != nil && nv.RangeVar != nil {
							result.tableWiths = append(result.tableWiths, TableWith{tableName: TableName(nv.RangeVar.Relname)})
						}
					}
				}
				if nj.JoinExpr.Rarg != nil && nj.JoinExpr.Rarg.GetNode() != nil {
					nr := nj.JoinExpr.Rarg.GetNode()
					nv, ok := nr.(*pg_query.Node_RangeVar)
					if ok {
						if nv != nil && nv.RangeVar != nil {
							result.tableWiths = append(result.tableWiths, TableWith{tableName: TableName(nv.RangeVar.Relname)})
						}
					}
				}
			}
		}
	}

	return result, nil
}

func parseInsertStmt(s *pg_query.InsertStmt) (*CRUDTableNames, error) {
	if s == nil {
		return nil, nil
	}

	result := &CRUDTableNames{crud: Insert}

	rel := s.GetRelation()
	if rel != nil {
		result.tableWiths = append(result.tableWiths, TableWith{tableName: TableName(rel.GetRelname())})
	}

	return result, nil
}

func parseUpdateStmt(s *pg_query.UpdateStmt) (*CRUDTableNames, error) {
	if s == nil {
		return nil, nil
	}

	result := &CRUDTableNames{crud: Update}

	rel := s.GetRelation()
	if rel != nil {
		result.tableWiths = append(result.tableWiths, TableWith{tableName: TableName(rel.GetRelname())})
	}

	return result, nil
}

func parseDeleteStmt(s *pg_query.DeleteStmt) (*CRUDTableNames, error) {
	if s == nil {
		return nil, nil
	}

	result := &CRUDTableNames{crud: Delete}

	rel := s.GetRelation()
	if rel != nil {
		result.tableWiths = append(result.tableWiths, TableWith{tableName: TableName(rel.GetRelname())})
	}

	return result, nil
}
