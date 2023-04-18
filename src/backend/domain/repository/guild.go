package repository

import "github.com/sky0621/familiagildo/domain/aggregate"

type GuildRepository interface {
	Save(a aggregate.GuildAggregate) error
}
