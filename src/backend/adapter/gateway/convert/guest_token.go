package convert

import (
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/vo"
	"github.com/sky0621/familiagildo/driver/db/generated"
)

func GuestTokenAggregateFromDBToDomain(src generated.GuestToken) *aggregate.GuestToken {
	return &aggregate.GuestToken{
		Root: GuestTokenFromDBToDomain(src),
		Guild: GuildFromDBToDomain(generated.Guild{
			ID: src.GuildID,
		}),
	}
}

func GuestTokenFromDBToDomain(src generated.GuestToken) *entity.GuestToken {
	return &entity.GuestToken{
		ID:             vo.ToID(src.ID),
		Token:          vo.ToToken(src.Token),
		ExpirationDate: vo.ToExpirationDate(src.ExpirationDate),
		Mail:           vo.ToOwnerMail(src.Mail),
	}
}
