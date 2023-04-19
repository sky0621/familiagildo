package repository

import (
	"context"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildRepository interface {
	CreateWithRegistering(ctx context.Context, name vo.GuildName) (*aggregate.GuildAggregate, error)
	UpdateWithRegistered(ctx context.Context, id vo.ID) (*aggregate.GuildAggregate, error)
}
