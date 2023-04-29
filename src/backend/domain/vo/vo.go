package vo

type ValueObject[T string | int64] interface {
	Validate() error
	FieldName() string
	ToVal() T
}
