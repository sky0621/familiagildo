package gateway

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *guildRepository) CreateWithRegistering(ctx context.Context, name vo.GuildName) (*aggregate.Guild, error) {
	q := r.db

	tx, ok := ctx.Value(app.TxCtxKey).(*sql.Tx)
	if ok {
		q = r.db.WithTx(tx)
	}

	record, err := q.CreateGuildWithRegistering(ctx, name.ToVal())
	if err != nil {
		return nil, app.WrapError(err, fmt.Sprintf("failed to CreateGuildWithRegistering [guildName:%s]", name.ToVal()))
	}
	return convert.GuildAggregateFromDBToDomain(record), nil
}

func (r *guildRepository) UpdateWithRegistered(ctx context.Context, id vo.ID) (*aggregate.Guild, error) {
	q := r.db

	tx, ok := ctx.Value(app.TxCtxKey).(*sql.Tx)
	if ok {
		q = r.db.WithTx(tx)
	}

	record, err := q.UpdateGuildWithRegistered(ctx, id.ToVal())
	if err != nil {
		return nil, app.WrapError(err, fmt.Sprintf("failed to UpdateGuildWithRegistered [id:%d]", id.ToVal()))
	}
	return convert.GuildAggregateFromDBToDomain(record), nil
}

func (r *guildRepository) GetByID(ctx context.Context, id vo.ID) (*aggregate.Guild, error) {
	q := r.db

	tx, ok := ctx.Value(app.TxCtxKey).(*sql.Tx)
	if ok {
		q = r.db.WithTx(tx)
	}

	record, err := q.GetGuildByID(ctx, id.ToVal())
	if err != nil {
		return nil, app.WrapError(err, fmt.Sprintf("failed to UpdateGuildWithRegistered [id:%d]", id.ToVal()))
	}
	return convert.GuildAggregateFromDBToDomain(record), nil
}
