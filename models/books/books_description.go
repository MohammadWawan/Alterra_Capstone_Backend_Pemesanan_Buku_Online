package books

import (
	"time"

	"gorm.io/gorm"
)

type Description struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Deskripsi string         `json:"deskripsi"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
