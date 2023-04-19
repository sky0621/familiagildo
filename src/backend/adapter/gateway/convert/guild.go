package convert

import (
	"github.com/sky0621/familiagildo/domain/aggregate"
	"github.com/sky0621/familiagildo/domain/entity"
	"github.com/sky0621/familiagildo/domain/vo"
	"github.com/sky0621/familiagildo/external/db"
)

func GuildAggregateFromDBToDomain(src db.Guild) *aggregate.GuildAggregate {
	return &aggregate.GuildAggregate{
		Guild: GuildFromDBToDomain(src),
		AuditItem: AuditFromDBToDomain(
			src.CreateUserID, src.UpdateUserID,
			src.DeleteUserID, src.CreatedAt,
			src.UpdatedAt, src.DeletedAt,
		),
	}
}

func GuildFromDBToDomain(src db.Guild) *entity.Guild {
	return &entity.Guild{
		ID:     vo.ParseID(src.ID),
		Name:   vo.ParseGuildName(src.Name),
		Status: vo.ParseGuildStatus(src.Status),
	}
}
