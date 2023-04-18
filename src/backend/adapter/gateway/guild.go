package gateway

import (
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/external/db"
)

func NewGuildRepository(db *db.Queries) repository.GuildRepository {
	return &guildRepository{db: db}
}

type guildRepository struct {
	db *db.Queries
}

func (r *guildRepository) Save(a aggregate.GuildAggregate) error {

	return nil
}
