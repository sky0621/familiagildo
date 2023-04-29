package vo

type ID int64

func (v ID) Validate() bool {
	if v <= 0 {
		return false
	}
	return true
}

func (v ID) FieldName() string {
	return "id"
}

func (v ID) ToVal() int64 {
	return int64(v)
}

func ParseID(v int64) ID {
	return ID(v)
}
