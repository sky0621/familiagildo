// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package generated

import (
	"database/sql"
	"time"
)

type Admin struct {
	ID       int64
	Name     string
	LoginID  sql.NullString
	Password sql.NullString
}

type GuestToken struct {
	ID             int64
	GuildID        int64
	Mail           string
	Token          string
	ExpirationDate time.Time
	AcceptedNumber int64
}

type Guild struct {
	ID           int64
	Name         string
	Status       int16
	CreateUserID int64
	CreatedAt    time.Time
	UpdateUserID sql.NullInt64
	UpdatedAt    sql.NullTime
	DeleteUserID sql.NullInt64
	DeletedAt    sql.NullTime
}

type GuildOwnerRelation struct {
	ID      int64
	GuildID int64
	OwnerID int64
}

type Owner struct {
	ID           int64
	Name         sql.NullString
	Mail         string
	LoginID      sql.NullString
	Password     sql.NullString
	CreateUserID sql.NullInt64
	CreatedAt    time.Time
	UpdateUserID sql.NullInt64
	UpdatedAt    sql.NullTime
	DeleteUserID sql.NullInt64
	DeletedAt    sql.NullTime
}

type Participant struct {
	ID           int64
	Name         string
	CreateUserID sql.NullInt64
	CreatedAt    time.Time
	UpdateUserID sql.NullInt64
	UpdatedAt    sql.NullTime
	DeleteUserID sql.NullInt64
	DeletedAt    sql.NullTime
}

type Task struct {
	ID           int64
	Content      string
	Status       int16
	Continuity   int16
	DueDatetime  sql.NullTime
	CreateUserID sql.NullInt64
	CreatedAt    time.Time
	UpdateUserID sql.NullInt64
	UpdatedAt    sql.NullTime
	DeleteUserID sql.NullInt64
	DeletedAt    sql.NullTime
}