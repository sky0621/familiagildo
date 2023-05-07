package vo

import "errors"

type GuildStatus int16

const (
	GuildStatusUndefined GuildStatus = iota
	GuildStatusRegistering
	GuildStatusRegistered
)

var UnMatchGuildStatusError = errors.New("UnMatchGuildStatusError")

func (v GuildStatus) Validate() error {
	if v != GuildStatusRegistering && v != GuildStatusRegistered {
		return UnMatchGuildStatusError
	}
	return nil
}

func (v GuildStatus) FieldName() string {
	return "guildStatus"
}

func (v GuildStatus) ToVal() int16 {
	return int16(v)
}

func ToGuildStatus(v int16) GuildStatus {
	switch v {
	case GuildStatusRegistering.ToVal():
		fallthrough
	case GuildStatusRegistered.ToVal():
		return GuildStatus(v)
	default:
		return GuildStatusUndefined
	}
}

func ParseGuildStatus(v int16) (GuildStatus, error) {
	gs := ToGuildStatus(v)
	if err := gs.Validate(); err != nil {
		return -1, err
	}
	return gs, nil
}
