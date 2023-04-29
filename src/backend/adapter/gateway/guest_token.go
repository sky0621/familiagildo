package gateway

import (
	"context"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/domain/vo"
	"github.com/sky0621/familiagildo/driver/db"
	"github.com/sky0621/familiagildo/driver/db/generated"
)

func NewGuestTokenRepository(cli *db.Client) repository.GuestTokenRepository {
	return &guestTokenRepository{db: cli.Q}
}

type guestTokenRepository struct {
	db *generated.Queries
}

func (r *guestTokenRepository) GetByOwnerMailWithinValidPeriod(ctx context.Context, mail vo.OwnerMail) (*aggregate.GuestToken, error) {
	// FIXME: implement me
	/*	return &aggregate.GuestToken{
			Root: &entity.GuestToken{
				Token: vo.ParseToken(uuid.New().String()),
			},
		}, nil
	*/
	return nil, nil
}
