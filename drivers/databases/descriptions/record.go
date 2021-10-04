package descriptions

import (
	"alterra/business/descriptions"
	"time"

	"gorm.io/gorm"
)

type Descriptions struct {
	Id          uint `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain descriptions.Domain) Descriptions {
	return Descriptions{
		Id:          domain.Id,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func (description *Descriptions) ToDomain() descriptions.Domain {
	return descriptions.Domain{
		Id:          description.Id,
		Description: description.Description,
		CreatedAt:   description.CreatedAt,
		UpdatedAt:   description.UpdatedAt,
	}
}

func ToListDomain(data []Descriptions) []descriptions.Domain {
	list := []descriptions.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
