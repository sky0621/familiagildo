package vo

type GuildName string

func (v GuildName) Validate() bool {
	if v == "" {
		return false
	}
	return true
}

func (v GuildName) ToVal() string {
	return string(v)
}

func ParseGuildName(v string) GuildName {
	return GuildName(v)
}
