package model

type Article struct {
	*Model
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
}

func (Article) TableName() string {
	return "tb_article"
}
