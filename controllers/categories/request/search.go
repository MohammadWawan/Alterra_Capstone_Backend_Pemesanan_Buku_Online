package request

import "alterra/business/categories"

type CategorySearch struct {
	Category string `json:"category"`
}

func ToDomain(search CategorySearch) categories.Domain {
	return categories.Domain{
		Category: search.Category,
	}
}
