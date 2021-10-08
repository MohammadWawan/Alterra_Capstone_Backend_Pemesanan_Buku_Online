package request

import (
	"alterra/business/categories"
)

type InsertCategory struct {
	Category string `json:"category"`
}

func (category *InsertCategory) ToDomain() *categories.Domain {
	return &categories.Domain{
		Category: category.Category,
	}
}
