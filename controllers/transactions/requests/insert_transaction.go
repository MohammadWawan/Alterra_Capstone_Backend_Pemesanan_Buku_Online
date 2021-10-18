package requests

import (
	"alterra/business/transactions"
)

type InsertTransaction struct {
	Method_Payment_Id uint `json:"method_payment_Id"`
	User_Id           uint `json:"user_id"`
	Karyawan_Id       uint `json:"karyawan_id"`
	Total_Qty         uint `json:"total_qty"`
	Total_Price       uint `json:"total_price"`
}

func (transaction *InsertTransaction) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		Method_Payment_Id: transaction.Method_Payment_Id,
		User_Id:           transaction.User_Id,
		Karyawan_Id:       transaction.Karyawan_Id,
		Total_Qty:         transaction.Total_Qty,
		Total_Price:       transaction.Total_Price,
	}
}
