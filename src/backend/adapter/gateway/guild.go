package gateway

import (
	"context"
	"github.com/sky0621/familiagildo/adapter/gateway/convert"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/domain/vo"
	"github.com/sky0621/familiagildo/external/db"
)

func NewGuildRepository(db *db.Queries) repository.GuildRepository {
	return &guildRepository{db: db}
}

type guildRepository struct {
	db *db.Queries
}

func (r *guildRepository) CreateWithRegistering(ctx context.Context, name vo.GuildName) (*aggregate.GuildAggregate, error) {
	record, err := r.db.CreateGuildWithRegistering(ctx, name.ToVal())
	if err != nil {
		return nil, app.WrapErrorWithMsgf(err, "name: %s", name)
	}
	return convert.GuildAggregateFromDBToDomain(record), nil
}

func (r *guildRepository) UpdateWithRegistered(ctx context.Context, id vo.ID) (*aggregate.GuildAggregate, error) {
	record, err := r.db.UpdateGuildWithRegistered(ctx, id.ToVal())
	if err != nil {
		return nil, app.WrapErrorWithMsgf(err, "id: %d", id)
	}
	return convert.GuildAggregateFromDBToDomain(record), nil
}
