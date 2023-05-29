package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (Tag) TableName() string {
	return "tb_tag"
}
