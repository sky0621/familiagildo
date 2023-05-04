package vo

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type OwnerName string

func (v OwnerName) Validate() error {
	if err := validation.Validate(v.ToVal(),
		validation.Required,
		validation.RuneLength(2, 100),
	); err != nil {
		return err
	}
	return nil
}

func (v OwnerName) FieldName() string {
	return "ownerName"
}

func (v OwnerName) ToVal() string {
	return string(v)
}

func ToOwnerName(v string) OwnerName {
	return OwnerName(v)
}
