package vo

type AcceptedNumber string

func (v AcceptedNumber) Validate() bool {
	if v == "" {
		return false
	}
	// FIXME:
	return true
}

func (v AcceptedNumber) FieldName() string {
	return "acceptedNumber"
}

func (v AcceptedNumber) ToVal() string {
	return string(v)
}

func ParseAcceptedNumber(v string) AcceptedNumber {
	return AcceptedNumber(v)
}
