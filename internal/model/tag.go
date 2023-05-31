package model

import "database/sql"

type Tag struct {
	*Model
	Name  string `json:"name"`
	Status sql.NullBool  `json:"status"`
}

func (Tag) TableName() string {
	return "tb_tag"
}
