package wishlist

import (
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	Id            int            `gorm:"primaryKey" json:"id"`
	User_Id       int            `json:"user_id" gorm:"unique"`
	Buku_Id       int            `json:"buku_id" gorm:"unique"`
	Wishlist_buku string         `json:"wishlist_buku"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
