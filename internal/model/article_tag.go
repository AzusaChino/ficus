package model

type ArticleTag struct {
	*Model
	TagId     int `json:"tag_id"`
	ArticleId int `json:"article_id"`
}

func (ArticleTag) TableName() string {
	return "tb_article_tag"
}
