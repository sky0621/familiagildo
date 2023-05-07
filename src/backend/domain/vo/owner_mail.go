package vo

import (
	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"
)

type OwnerMail string

func (v OwnerMail) Validate() error {
	if err := validation.Validate(v.ToVal(),
		validation.Required,
		is.Email,
		// FIXME:
	); err != nil {
		return err
	}
	return nil
}

func (v OwnerMail) FieldName() string {
	return "ownerMail"
}

func (v OwnerMail) ToVal() string {
	return string(v)
}

func ToOwnerMail(v string) OwnerMail {
	return OwnerMail(v)
}
