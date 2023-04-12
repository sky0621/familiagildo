package vo

type ValueObject interface {
	Validate() bool
}
