package requests

import "alterra/business/payment_methods"

type InsertPayment_Method struct {
	Type string `json:"type"`
}

func (payment_method *InsertPayment_Method) ToDomain() *payment_methods.Domain {
	return &payment_methods.Domain{
		Type: payment_method.Type,
	}
}
