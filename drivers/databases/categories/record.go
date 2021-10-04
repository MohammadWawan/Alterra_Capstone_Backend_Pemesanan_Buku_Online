package categories

import (
	"alterra/business/categories"
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	Id            uint `gorm:"primaryKey"`
	Name_Category string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain categories.Domain) Categories {
	return Categories{
		Id:            domain.Id,
		Name_Category: domain.Name_Category,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func (category *Categories) ToDomain() categories.Domain {
	return categories.Domain{
		Id:            category.Id,
		Name_Category: category.Name_Category,
		CreatedAt:     category.CreatedAt,
		UpdatedAt:     category.UpdatedAt,
	}
}

func ToListDomain(data []Categories) []categories.Domain {
	list := []categories.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
