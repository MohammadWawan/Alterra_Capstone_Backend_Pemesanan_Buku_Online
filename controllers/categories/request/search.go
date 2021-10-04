package request

import "alterra/business/categories"

type CategorySearch struct {
	Name_Category string `json:"category"`
}

func ToDomain(search CategorySearch) categories.Domain {
	return categories.Domain{
		Name_Category: search.Name_Category,
	}
}
