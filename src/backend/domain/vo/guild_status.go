package vo

type GuildStatus int16

const (
	GuildStatusUndefined GuildStatus = iota
	GuildStatusRegistering
	GuildStatusRegistered
)

func (v GuildStatus) Validate() bool {
	if v == 0 {
		return false
	}
	if v != GuildStatusRegistering && v != GuildStatusRegistered {
		return false
	}
	return true
}

func (v GuildStatus) FieldName() string {
	return "guildStatus"
}

func (v GuildStatus) ToVal() int16 {
	return int16(v)
}

func ParseGuildStatus(v int16) GuildStatus {
	switch v {
	case GuildStatusRegistering.ToVal():
		fallthrough
	case GuildStatusRegistered.ToVal():
		return GuildStatus(v)
	default:
		return GuildStatusUndefined
	}
}
