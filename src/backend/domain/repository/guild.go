package repository

import (
	"context"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildRepository interface {
	// CreateWithRegistering is ギルドの仮登録
	CreateWithRegistering(ctx context.Context, name vo.GuildName) (*aggregate.GuildAggregate, error)
	// UpdateWithRegistered is ギルドの本登録
	UpdateWithRegistered(ctx context.Context, id vo.ID) (*aggregate.GuildAggregate, error)
}
