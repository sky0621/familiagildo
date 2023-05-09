package aggregate

import "github.com/sky0621/familiagildo/domain/entity"

type Guild struct {
	Root      *entity.Guild
	Owner     *entity.Owner
	AuditItem *entity.AuditItem
}
