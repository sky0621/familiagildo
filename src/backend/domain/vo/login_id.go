package vo

import validation "github.com/go-ozzo/ozzo-validation"

type LoginID string

func (v LoginID) Validate() error {
	if err := validation.Validate(v.ToVal(),
		validation.Required,
		// FIXME:
	); err != nil {
		return err
	}
	return nil
}

func (v LoginID) FieldName() string {
	return "id"
}

func (v LoginID) ToVal() string {
	return string(v)
}

func ParseLoginID(v string) LoginID {
	return LoginID(v)
}
