package transaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaksi_Detail struct {
	Buku_Id      int            `gorm:"primaryKey" json:"buku_id"`
	Transaksi_Id int            `gorm:"primaryKey" json:"transaksi_id"`
	Qty          int            `json:"qty_id"`
	Harga        int            `json:"harga"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
