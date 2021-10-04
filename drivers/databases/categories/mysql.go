package categories

import (
	"alterra/business/categories"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MysqlCategoryRepository struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository(conn *gorm.DB) *MysqlCategoryRepository {
	return &MysqlCategoryRepository{
		Conn: conn,
	}
}

func (rep *MysqlCategoryRepository) GetListCategory(ctx context.Context, search string) ([]categories.Domain, error) {
	var data []Categories
	err := rep.Conn.Find(&data)
	if err.Error != nil {
		return []categories.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *MysqlCategoryRepository) GetById(ctx context.Context, Id uint) (categories.Domain, error) {
	var category Categories
	result := rep.Conn.Find(&category, "id = ?", Id)
	if result.Error != nil {
		return categories.Domain{}, result.Error
	}

	return category.ToDomain(), nil
}

func (rep *MysqlCategoryRepository) InsertCategory(_ context.Context, domain categories.Domain) (categories.Domain, error) {
	category := FromDomain(domain)
	err := rep.Conn.Create(&category)
	if err.Error != nil {
		return categories.Domain{}, err.Error
	}
	return category.ToDomain(), nil
}

func (rep *MysqlCategoryRepository) Update(ctx context.Context, domain categories.Domain, id uint) (categories.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return categories.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *MysqlCategoryRepository) Delete(ctx context.Context, id uint) error {
	category := Categories{}
	err := rep.Conn.Delete(&category, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
