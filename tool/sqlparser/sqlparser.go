package sqlparser

import (
	"errors"
	"fmt"
	"log"

	pg_query "github.com/pganalyze/pg_query_go/v4"
)

// sql-name
//   table-name
//     crud(s)

type SQLParser interface {
	Parse(sqlName, sql string) (*SQLParseResult, error)
}

func NewSQLParser() SQLParser {
	return &sqlParser{}
}

type sqlParser struct {
}

func (p *sqlParser) Parse(sqlName, sql string) (*SQLParseResult, error) {
	res, err := pg_query.Parse(sql)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("result == nil")
	}

	processedResult := &SQLParseResult{sqlName: ToSQLName(sqlName)}

	fmt.Println("===========================")
	for _, stmt := range res.GetStmts() {
		tableNameWithCRUDSlice, err := parseStmt(stmt.GetStmt())
		if err != nil {
			log.Println("parseStmt is nil")
			continue
		}
		if tableNameWithCRUDSlice == nil || len(tableNameWithCRUDSlice) == 0 {
			continue
		}
		//		processedResult.tableNameWithCRUDsSlice = append(processedResult.tableNameWithCRUDsSlice, tableNameWithCRUDs)
		for _, x := range tableNameWithCRUDSlice {
			fmt.Println(x.tableName)
			fmt.Println(x.crud)
		}
	}
	fmt.Println("===========================")

	/*	fmt.Println("===========================")
		fmt.Println(processedResult.sqlName)
		for _, tc := range processedResult.tableNameWithCRUDsSlice {
			fmt.Println(tc.tableName)
			fmt.Println(tc.cruds)
		}
	*/

	return processedResult, nil
}

type SQLParseResult struct {
	sqlName                 SQLName
	tableNameWithCRUDsSlice []*TableNameWithCRUDs
}

type TableNameWithCRUDs struct {
	tableName TableName
	cruds     []CRUD
}

type TableNameWithCRUD struct {
	tableName TableName
	crud      CRUD
}

type SQLName string

func ToSQLName(n string) SQLName {
	return SQLName(n)
}

type TableName string

type CRUD int8

const (
	Select CRUD = iota + 1
	Insert
	Update
	Delete
)

func parseStmt(s *pg_query.Node) ([]*TableNameWithCRUD, error) {
	if s == nil {
		return nil, errors.New("node is nil")
	}

	var tableNameWithCRUDSlice []*TableNameWithCRUD
	var err error

	tableNameWithCRUDSlice, err = parseSelectStmt(s.GetSelectStmt())
	if err != nil {
		return nil, err
	}
	if tableNameWithCRUDSlice != nil {
		return tableNameWithCRUDSlice, nil
	}

	tableNameWithCRUDSlice, err = parseInsertStmt(s.GetInsertStmt())
	if err != nil {
		return nil, err
	}
	if tableNameWithCRUDSlice != nil {
		return tableNameWithCRUDSlice, nil
	}

	tableNameWithCRUDSlice, err = parseUpdateStmt(s.GetUpdateStmt())
	if err != nil {
		return nil, err
	}
	if tableNameWithCRUDSlice != nil {
		return tableNameWithCRUDSlice, nil
	}

	tableNameWithCRUDSlice, err = parseDeleteStmt(s.GetDeleteStmt())
	if err != nil {
		return nil, err
	}
	if tableNameWithCRUDSlice != nil {
		return tableNameWithCRUDSlice, nil
	}

	return nil, nil
}

func parseSelectStmt(s *pg_query.SelectStmt) ([]*TableNameWithCRUD, error) {
	if s == nil {
		return nil, nil
	}

	fromArray := s.FromClause
	if fromArray == nil {
		return nil, nil
	}

	var result []*TableNameWithCRUD
	crud := Select

	for _, from := range fromArray {
		n := from.GetNode()
		nv, ok := n.(*pg_query.Node_RangeVar)
		if ok {
			if nv != nil && nv.RangeVar != nil {
				result = append(result, &TableNameWithCRUD{crud: crud, tableName: TableName(nv.RangeVar.Relname)})
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
							result = append(result, &TableNameWithCRUD{crud: crud, tableName: TableName(nv.RangeVar.Relname)})
						}
					}
				}
				if nj.JoinExpr.Rarg != nil && nj.JoinExpr.Rarg.GetNode() != nil {
					nr := nj.JoinExpr.Rarg.GetNode()
					nv, ok := nr.(*pg_query.Node_RangeVar)
					if ok {
						if nv != nil && nv.RangeVar != nil {
							result = append(result, &TableNameWithCRUD{crud: crud, tableName: TableName(nv.RangeVar.Relname)})
						}
					}
				}
			}
		}
	}

	return result, nil
}

func parseInsertStmt(s *pg_query.InsertStmt) ([]*TableNameWithCRUD, error) {
	if s == nil {
		return nil, nil
	}

	var result []*TableNameWithCRUD
	crud := Insert

	rel := s.GetRelation()
	if rel != nil {
		result = append(result, &TableNameWithCRUD{crud: crud, tableName: TableName(rel.GetRelname())})
	}

	return result, nil
}

func parseUpdateStmt(s *pg_query.UpdateStmt) ([]*TableNameWithCRUD, error) {
	if s == nil {
		return nil, nil
	}

	var result []*TableNameWithCRUD
	crud := Update

	rel := s.GetRelation()
	if rel != nil {
		result = append(result, &TableNameWithCRUD{crud: crud, tableName: TableName(rel.GetRelname())})
	}

	return result, nil
}

func parseDeleteStmt(s *pg_query.DeleteStmt) ([]*TableNameWithCRUD, error) {
	if s == nil {
		return nil, nil
	}

	var result []*TableNameWithCRUD
	crud := Delete

	rel := s.GetRelation()
	if rel != nil {
		result = append(result, &TableNameWithCRUD{crud: crud, tableName: TableName(rel.GetRelname())})
	}

	return result, nil
}
