package models

import (
	"time"

	"gorm.io/gorm"
)

//Model ...
type Model struct {
	ID        uint            `json:"id"`
	CreatedAt time.Time       `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time       `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
