package sqlparser

// sql-name
//   table-name
//     crud(s)

type SQLParser interface {
	Parse(sql string) (*SQLParseResult, error)
}

func NewSQLParser() SQLParser {
	return &sqlParser{}
}

type sqlParser struct {
}

func (p *sqlParser) Parse(sql string) (*SQLParseResult, error) {
	// FIXME:
	return nil, nil
}

type SQLParseResult struct {
	SQLName                 SQLName
	TableNameWithCRUDsSlice []TableNameWithCRUDs
}

type TableNameWithCRUDs struct {
	TableName TableName
	CRUDs     []CRUD
}

type SQLName string

type TableName string

type CRUD int8

const (
	Select CRUD = iota + 1
	Insert
	Update
	Delete
)
