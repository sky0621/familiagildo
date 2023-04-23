package vo

type ValueObject interface {
	Validate() error
}
