package books

import (
	"time"

	"gorm.io/gorm"
)

type Deskripsi struct {
	Id          int            `gorm:"primaryKey" json:"id"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
