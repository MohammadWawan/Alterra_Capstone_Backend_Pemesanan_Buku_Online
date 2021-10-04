package books

import (
	"time"

	"gorm.io/gorm"
)

type Kategori struct {
	Id         int            `gorm:"primaryKey" json:"id"`
	Categories string         `json:"categories"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
