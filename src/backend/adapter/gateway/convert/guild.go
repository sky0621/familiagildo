package convert

import (
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/vo"
	"github.com/sky0621/familiagildo/driver/db/generated"
)

func GuildAggregateFromDBToDomain(src generated.Guild) *aggregate.GuildAggregate {
	return &aggregate.GuildAggregate{
		Guild: GuildFromDBToDomain(src),
		AuditItem: AuditFromDBToDomain(
			src.CreateUserID, src.UpdateUserID,
			src.DeleteUserID, src.CreatedAt,
			src.UpdatedAt, src.DeletedAt,
		),
	}
}

func GuildFromDBToDomain(src generated.Guild) *entity.Guild {
	return &entity.Guild{
		ID:     vo.ParseID(src.ID),
		Name:   vo.ParseGuildName(src.Name),
		Status: vo.ParseGuildStatus(src.Status),
	}
}
