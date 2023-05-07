package vo

type ID int64

func (v ID) Validate() error {
	// FIXME:
	return nil
}

func (v ID) FieldName() string {
	return "id"
}

func (v ID) ToVal() int64 {
	return int64(v)
}

func ToID(v int64) ID {
	return ID(v)
}

func ParseID(v int64) (ID, error) {
	id := ToID(v)
	if err := id.Validate(); err != nil {
		return -1, err
	}
	return id, nil
}
