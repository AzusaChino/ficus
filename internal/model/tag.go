package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name   string       `json:"name"`
	Status sql.NullBool `json:"status"`
}

func (Tag) TableName() string {
	return "tb_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("status = ?", t.Status)
	if err := db.Model(&t).Where("is_deleted = ?", false).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (t Tag) Get(db *gorm.DB) (*Tag, error) {
	var tag Tag
	db = db.Where("id = ?", t.ID)
	if err := db.Where("is_deleted = ?", false).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}
