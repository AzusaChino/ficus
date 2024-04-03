package model

import (
	"database/sql"
	"time"
)

type Model struct {
	ID         int          `gorm:"primary_key" json:"id"`
	CreatedBy  string       `json:"created_by"`
	CreatedAt  time.Time    `json:"created_at"`
	ModifiedBy string       `json:"modified_by"`
	ModifiedAt time.Time    `json:"modified_at"`
	IsDeleted  sql.NullBool `json:"is_deleted"`
	DeletedAt  time.Time    `json:"deleted_at"`
}
