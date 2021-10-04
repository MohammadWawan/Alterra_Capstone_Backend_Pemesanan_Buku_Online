package books

import (
	"alterra/business/books"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MysqlBookRepository struct {
	Conn *gorm.DB
}

func NewMysqlBookRepository(conn *gorm.DB) *MysqlBookRepository {
	return &MysqlBookRepository{
		Conn: conn,
	}
}

func (rep *MysqlBookRepository) GetListBook(ctx context.Context, search string) ([]books.Domain, error) {
	var data []Books
	err := rep.Conn.Find(&data)
	if err.Error != nil {
		return []books.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *MysqlBookRepository) GetById(ctx context.Context, Id uint) (books.Domain, error) {
	var book Books
	result := rep.Conn.Find(&book, "id = ?", Id)
	if result.Error != nil {
		return books.Domain{}, result.Error
	}

	return book.ToDomain(), nil
}

func (rep *MysqlBookRepository) InsertBook(_ context.Context, domain books.Domain) (books.Domain, error) {
	book := FromDomain(domain)
	err := rep.Conn.Create(&book)
	if err.Error != nil {
		return books.Domain{}, err.Error
	}
	return book.ToDomain(), nil
}

func (rep *MysqlBookRepository) Update(ctx context.Context, domain books.Domain, id uint) (books.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return books.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *MysqlBookRepository) Delete(ctx context.Context, id uint) error {
	book := Books{}
	err := rep.Conn.Delete(&book, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
