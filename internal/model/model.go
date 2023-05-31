package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/azusachino/ficus/pkg/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func NewDbEngine(dbConfig *conf.DatabaseConfig) (*gorm.DB, error) {
	// open postgres connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbName, dbConfig.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}
