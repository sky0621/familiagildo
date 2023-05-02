package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sky0621/familiagildo/adapter/gateway/convert"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
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
	record, err := r.db.GetGuestTokenByMailWithinExpirationDate(ctx, mail.ToVal())
	if err != nil {
		if IsNoRecords(err) {
			return nil, nil
		}
		return nil, app.WrapError(err, fmt.Sprintf("failed to GetByOwnerMailWithinValidPeriod [mail:%s]", mail.ToVal()))
	}
	return convert.GuestTokenAggregateFromDBToDomain(record), nil
}

func (r *guestTokenRepository) Create(ctx context.Context, guildID vo.ID, mail vo.OwnerMail, guestToken *entity.GuestToken, acceptedNumber vo.AcceptedNumber) (*aggregate.GuestToken, error) {
	q := r.db

	tx, ok := ctx.Value(app.TxCtxKey).(*sql.Tx)
	if ok {
		q = r.db.WithTx(tx)
	}

	record, err := q.CreateGuestToken(ctx, generated.CreateGuestTokenParams{
		GuildID:        guildID.ToVal(),
		Mail:           mail.ToVal(),
		Token:          guestToken.Token.ToVal(),
		ExpirationDate: guestToken.ExpirationDate.ToVal(),
		AcceptedNumber: acceptedNumber.ToVal(),
	})
	if err != nil {
		return nil, app.WrapError(err, fmt.Sprintf("failed to CreateGuestToken [guildID:%d][mail:%s][token:%s][expirationDate:%v][acceptedNumber:%s]",
			guildID.ToVal(), mail.ToVal(), guestToken.Token.ToVal(), guestToken.ExpirationDate.ToVal(), acceptedNumber.ToVal()))
	}
	return convert.GuestTokenAggregateFromDBToDomain(record), nil
}
