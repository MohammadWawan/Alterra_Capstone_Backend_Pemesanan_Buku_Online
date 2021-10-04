package books

import (
	"alterra/business/books"
	"alterra/drivers/databases/categories"
	"alterra/drivers/databases/descriptions"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Id             uint `gorm:"primaryKey"`
	Category_Id    uint
	Description_Id uint
	Title          string
	Price          uint
	Author         string
	Publisher      string
	Category       categories.Categories     `gorm:"foreignKey:kategori_id"`
	Description    descriptions.Descriptions `gorm:"foreignkey:Description;references:description_id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain books.Domain) Books {
	return Books{
		Id:          domain.Id,
		Category_Id: domain.Category_Id,
		Title:       domain.Title,
		Price:       domain.Price,
		Author:      domain.Author,
		Category:    categories.Categories{},
		Description: descriptions.Descriptions{},
		Publisher:   domain.Publisher,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func (book *Books) ToDomain() books.Domain {
	return books.Domain{
		Id:             book.Id,
		Category_Id:    book.Category_Id,
		Description_Id: book.Description_Id,
		Title:          book.Title,
		Price:          book.Price,
		Author:         book.Author,
		Publisher:      book.Publisher,
		CreatedAt:      book.CreatedAt,
		UpdatedAt:      book.UpdatedAt,
	}
}

func ToListDomain(data []Books) []books.Domain {
	list := []books.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
