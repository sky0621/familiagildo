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

	for _, stmt := range res.GetStmts() {
		result, err := parseStmt(stmt.GetStmt())
		if err != nil {
			log.Println("parseStmt is nil")
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

	return processedResult, nil
}

type SQLParseResult struct {
	sqlName                 SQLName
	tableNameWithCRUDsSlice []TableNameWithCRUDs
}

type TableNameWithCRUDs struct {
	tableName TableName
	cruds     []CRUD
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
