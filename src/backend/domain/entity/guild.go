package entity

import "github.com/sky0621/familiagildo/domain/vo"

type Guild struct {
	ID     vo.ID
	Name   vo.GuildName
	Status vo.GuildStatus
}
