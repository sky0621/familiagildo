package vo

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type GuildName string

func (v GuildName) Validate() error {
	if err := validation.Validate(v.ToVal(),
		validation.Required,
		validation.RuneLength(2, 100),
		// FIXME:
	); err != nil {
		return err
	}
	return nil
}

func (v GuildName) FieldName() string {
	return "guildName"
}

func (v GuildName) ToVal() string {
	return string(v)
}

func ToGuildName(v string) GuildName {
	return GuildName(v)
}

func ParseGuildName(v string) (GuildName, error) {
	gn := ToGuildName(v)
	if err := gn.Validate(); err != nil {
		return "", err
	}
	return gn, nil
}
