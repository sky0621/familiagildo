package gateway

import (
	"context"
	"github.com/sky0621/familiagildo/adapter/gateway/convert"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/domain/vo"
	"github.com/sky0621/familiagildo/driver/db"
	"github.com/sky0621/familiagildo/driver/db/generated"
)

func NewGuildRepository(cli *db.Client) repository.GuildRepository {
	return &guildRepository{db: cli.Q}
}

type guildRepository struct {
	db *generated.Queries
}

func (r *guildRepository) CreateWithRegistering(ctx context.Context, name vo.GuildName) (*aggregate.GuildAggregate, error) {
	record, err := r.db.CreateGuildWithRegistering(ctx, name.ToVal())
	if err != nil {
		return nil, &app.AuthenticationError{}
	}
	return convert.GuildAggregateFromDBToDomain(record), nil
}

func (r *guildRepository) UpdateWithRegistered(ctx context.Context, id vo.ID) (*aggregate.GuildAggregate, error) {
	record, err := r.db.UpdateGuildWithRegistered(ctx, id.ToVal())
	if err != nil {
		return nil, app.WrapErrorWithMsgf(err, app.UnexpectedFailure, "id: %d", id)
	}
	return convert.GuildAggregateFromDBToDomain(record), nil
}
