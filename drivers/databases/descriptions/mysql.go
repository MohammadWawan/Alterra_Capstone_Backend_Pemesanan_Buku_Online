package descriptions

import (
	"alterra/business/descriptions"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MysqlDescriptionRepository struct {
	Conn *gorm.DB
}

func NewMysqlDescriptionRepository(conn *gorm.DB) *MysqlDescriptionRepository {
	return &MysqlDescriptionRepository{
		Conn: conn,
	}
}

func (rep *MysqlDescriptionRepository) GetListDescription(ctx context.Context, search string) ([]descriptions.Domain, error) {
	var data []Descriptions
	err := rep.Conn.Find(&data)
	if err.Error != nil {
		return []descriptions.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *MysqlDescriptionRepository) GetById(ctx context.Context, Id uint) (descriptions.Domain, error) {
	var description Descriptions
	result := rep.Conn.Find(&description, "id = ?", Id)
	if result.Error != nil {
		return descriptions.Domain{}, result.Error
	}

	return description.ToDomain(), nil
}

func (rep *MysqlDescriptionRepository) InsertDescription(ctx context.Context, domain descriptions.Domain) (descriptions.Domain, error) {
	description := FromDomain(domain)
	err := rep.Conn.Create(&description)
	if err.Error != nil {
		return descriptions.Domain{}, err.Error
	}
	return description.ToDomain(), nil
}

func (rep *MysqlDescriptionRepository) Update(ctx context.Context, domain descriptions.Domain, id uint) (descriptions.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return descriptions.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *MysqlDescriptionRepository) Delete(ctx context.Context, id uint) error {
	description := Descriptions{}
	err := rep.Conn.Delete(&description, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
