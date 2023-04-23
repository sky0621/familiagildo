package vo

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type GuildName string

func (v GuildName) Validate() error {
	if err := validation.Validate(v.ToVal(),
		validation.Required,
		validation.RuneLength(2, 100),
	); err != nil {
		return err
	}
	return nil
}

func (v GuildName) ToVal() string {
	return string(v)
}

func ToGuildName(v string) GuildName {
	return GuildName(v)
}
