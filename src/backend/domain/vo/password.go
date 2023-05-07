package vo

import validation "github.com/go-ozzo/ozzo-validation"

type Password string

func (v Password) Validate() error {
	if err := validation.Validate(v.ToVal(),
		validation.Required,
		// FIXME:
	); err != nil {
		return err
	}
	return nil
}

func (v Password) FieldName() string {
	return "password"
}

func (v Password) ToVal() string {
	return string(v)
}

func ToPassword(v string) Password {
	return Password(v)
}

func ParsePassword(v string) (Password, error) {
	li := ToPassword(v)
	if err := li.Validate(); err != nil {
		return "", err
	}
	return li, nil
}
