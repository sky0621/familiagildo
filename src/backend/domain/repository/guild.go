package repository

import (
	"context"

	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/vo"
)

type GuildRepository interface {
	CreateWithRegistering(ctx context.Context, name vo.GuildName) (*aggregate.Guild, error) // 仮登録
	UpdateWithRegistered(ctx context.Context, id vo.ID) (*aggregate.Guild, error)           // 本登録
	GetByID(ctx context.Context, id vo.ID) (*aggregate.Guild, error)
}
