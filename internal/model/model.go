package model

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  string `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn string `json:"modified_on"`
	DeletedOn  string `json:"deleted_on"`
	IsDel      uint8   `json:"is_del"`
}
