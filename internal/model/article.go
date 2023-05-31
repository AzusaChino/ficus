package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Brief         string `json:"brief"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
}

func (Article) TableName() string {
	return "tb_article"
}
