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

func ToLoginID(v string) LoginID {
	return LoginID(v)
}

func ParseLoginID(v string) (LoginID, error) {
	li := ToLoginID(v)
	if err := li.Validate(); err != nil {
		return "", err
	}
	return li, nil
}
