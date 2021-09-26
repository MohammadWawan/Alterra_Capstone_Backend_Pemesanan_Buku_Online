package books

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Id          int            `gorm:"primaryKey" json:"id"`
	Kategori_id int            `json:"kategori_id" gorm:"unique"`
	Title       string         `json:"title"`
	Price       int            `json:"price"`
	Author      string         `json:"author"`
	Publisher   string         `json:"publisher"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
