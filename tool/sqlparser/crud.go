package sqlparser

type CRUD int8

const (
	Create CRUD = iota + 1
	Read
	Update
	Delete
)

func (c CRUD) ToName() string {
	switch c {
	case Create:
		return "CREATE"
	case Read:
		return "READ"
	case Update:
		return "UPDATE"
	case Delete:
		return "DELETE"
	}
	return ""
}

func (c CRUD) ToShortName() string {
	switch c {
	case Create:
		return "C"
	case Read:
		return "R"
	case Update:
		return "U"
	case Delete:
		return "D"
	}
	return ""
}
