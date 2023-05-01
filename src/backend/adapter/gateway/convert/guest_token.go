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
	}
}

func GuestTokenFromDBToDomain(src generated.GuestToken) *entity.GuestToken {
	return &entity.GuestToken{
		ID:             vo.ParseID(src.ID),
		Token:          vo.ParseToken(src.Token),
		ExpirationDate: vo.ParseExpirationDate(src.ExpirationDate),
	}
}
