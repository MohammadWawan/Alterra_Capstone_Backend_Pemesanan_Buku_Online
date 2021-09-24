package methodPayments

import (
	"time"

	"gorm.io/gorm"
)

type Method_Payment struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Jenis     string         `json:"jenis"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
