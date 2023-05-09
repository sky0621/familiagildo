package entity

import "github.com/sky0621/familiagildo/domain/vo"

type Owner struct {
	ID       vo.ID
	Name     vo.OwnerName
	Mail     vo.OwnerMail
	LoginID  vo.LoginID
	Password vo.Password
}
