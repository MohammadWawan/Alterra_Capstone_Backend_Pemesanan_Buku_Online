package books

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Kategori  string         `json:"kategori"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
