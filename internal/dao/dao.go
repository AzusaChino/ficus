package dao

import (
	"database/sql"

	"github.com/azusachino/ficus/internal/model"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d *Dao) CountTag(name string, status bool) (int, error) {
	tag := model.Tag{Name: name, Status: sql.NullBool{Bool: status, Valid: true}}
	return tag.Count(d.engine)
}

func (d *Dao) GetTag(id int) (*model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Get(d.engine)
}
