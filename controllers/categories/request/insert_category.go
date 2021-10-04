package request

import (
	"alterra/business/categories"
)

type InsertCategory struct {
	Name_Category string `json:"category"`
}

func (category *InsertCategory) ToDomain() *categories.Domain {
	return &categories.Domain{
		Name_Category: category.Name_Category,
	}
}
