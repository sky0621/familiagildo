package vo

import validation "github.com/go-ozzo/ozzo-validation"

type AcceptedNumber string

func (v AcceptedNumber) Validate() error {
	if err := validation.Validate(v.ToVal(),
		validation.Required,
		// FIXME:
	); err != nil {
		return err
	}
	return nil
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
