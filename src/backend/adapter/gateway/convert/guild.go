package convert

import (
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/vo"
	"github.com/sky0621/familiagildo/driver/db/generated"
)

func GuildAggregateFromDBToDomain(src generated.Guild) *aggregate.Guild {
	return &aggregate.Guild{
		Root: GuildFromDBToDomain(src),
		AuditItem: AuditFromDBToDomain(
			src.CreateUserID, src.UpdateUserID,
			src.DeleteUserID, src.CreatedAt,
			src.UpdatedAt, src.DeletedAt,
		),
	}
}

func GuildFromDBToDomain(src generated.Guild) *entity.Guild {
	return &entity.Guild{
		ID:     vo.ToID(src.ID),
		Name:   vo.ToGuildName(src.Name),
		Status: vo.ToGuildStatus(src.Status),
	}
}
