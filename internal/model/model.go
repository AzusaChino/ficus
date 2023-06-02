package model

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/azusachino/ficus/global"
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
	pgHost := os.Getenv(global.PG_HOST)
	pgPass := os.Getenv(global.PG_PASS)
	// open postgres connection
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbConfig.DbUser, pgPass, pgHost, dbConfig.DbPort, dbConfig.DbName, dbConfig.SslMode)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}
