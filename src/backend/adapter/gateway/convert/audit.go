package convert

import (
	"database/sql"
	"github.com/sky0621/familiagildo/domain/entity"
	"time"
)

func NullableIDFromDBToDomain(src sql.NullInt64) *int64 {
	if !src.Valid {
		return nil
	}
	return &src.Int64
}

func NullableIDFromDomainToDB(src *int64) sql.NullInt64 {
	if src == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: *src,
		Valid: true,
	}
}

func NullableDateTimeFromDBToDomain(src sql.NullTime) *time.Time {
	if !src.Valid {
		return nil
	}
	return &src.Time
}

func NullableDateTimeFromDomainToDB(src *time.Time) sql.NullTime {
	if src == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  *src,
		Valid: true,
	}
}

func AuditFromDBToDomain(createUserID int64, updateUserID, deleteUserID sql.NullInt64, createdAt time.Time, updatedAt, deletedAt sql.NullTime) *entity.AuditItem {
	return &entity.AuditItem{
		CreateUserID: createUserID,
		CreatedAt:    createdAt,
		UpdateUserID: NullableIDFromDBToDomain(updateUserID),
		UpdatedAt:    NullableDateTimeFromDBToDomain(updatedAt),
		DeleteUserID: NullableIDFromDBToDomain(deleteUserID),
		DeletedAt:    NullableDateTimeFromDBToDomain(deletedAt),
	}
}
