package transaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaksi struct {
	Id                   int            `gorm:"primaryKey" json:"id"`
	Metode_Pembayaran_id int            `json:"kategori_id" gorm:"unique"`
	User_id              int            `json:"user_id" gorm:"unique"`
	Karyawan_id          int            `json:"karyawan_id" gorm:"unique"`
	Total_Qty            int            `json:"total_qty_id"`
	Total_Bayar          int            `json:"total_bayar_id"`
	CreatedAt            time.Time      `json:"createdAt"`
	UpdatedAt            time.Time      `json:"updatedAt"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
