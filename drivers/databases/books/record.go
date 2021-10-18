package books

import (
	"alterra/business/books"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Id             uint `gorm:"primaryKey"`
	Title          string
	Price          uint
	Author         string
	Publisher      string
	Category_Id    uint
	Description_Id uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func FromDomain(domain books.Domain) Books {
	return Books{
		Id:             domain.Id,
		Title:          domain.Title,
		Price:          domain.Price,
		Author:         domain.Author,
		Publisher:      domain.Publisher,
		Category_Id:    domain.Category_Id,
		Description_Id: domain.Category_Id,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}

func (book *Books) ToDomain() books.Domain {
	return books.Domain{
		Id:             book.Id,
		Title:          book.Title,
		Price:          book.Price,
		Author:         book.Author,
		Publisher:      book.Publisher,
		Category_Id:    book.Category_Id,
		Description_Id: book.Description_Id,
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
