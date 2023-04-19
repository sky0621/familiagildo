package entity

import "time"

type AuditItem struct {
	CreateUserID int64
	CreatedAt    time.Time
	UpdateUserID *int64
	UpdatedAt    *time.Time
	DeleteUserID *int64
	DeletedAt    *time.Time
}
