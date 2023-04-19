package aggregate

import "github.com/sky0621/familiagildo/domain/entity"

type GuildAggregate struct {
	Guild     *entity.Guild
	AuditItem *entity.AuditItem
}
