package requests

import "alterra/business/payment_methods"

type Payment_MethodSearch struct {
	Type string `json:"type"`
}

func ToDomain(search Payment_MethodSearch) payment_methods.Domain {
	return payment_methods.Domain{
		Type: search.Type,
	}
}
