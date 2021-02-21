package models

//Category 分类
type Category struct {
	CategoryID   uint   `json:"categoryID" db:"category_id"`
	CategoryName string `json:"categoryName" db:"category_name"`
}
