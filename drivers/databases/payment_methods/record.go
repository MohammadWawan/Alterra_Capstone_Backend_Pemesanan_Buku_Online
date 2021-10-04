package payment_methods

import (
	"alterra/business/payment_methods"
	"time"

	"gorm.io/gorm"
)

type Payment_Method struct {
	Id        uint `gorm:"primaryKey"`
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain payment_methods.Domain) Payment_Method {
	return Payment_Method{
		Id:        domain.Id,
		Type:      domain.Type,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func (payment_method *Payment_Method) ToDomain() payment_methods.Domain {
	return payment_methods.Domain{
		Id:        payment_method.Id,
		Type:      payment_method.Type,
		CreatedAt: payment_method.CreatedAt,
		UpdatedAt: payment_method.UpdatedAt,
	}
}

func ToListDomain(data []Payment_Method) []payment_methods.Domain {
	list := []payment_methods.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
