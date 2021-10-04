package users

import (
	"alterra/business/users"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) *MysqlUserRepository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	var user Users
	err := rep.Conn.Find(&user, "email = ?", email)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Register(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	user := Users{
		Id:       domain.Id,
		Name:     domain.Name,
		Email:    domain.Email,
		Address:  domain.Address,
		Password: domain.Password,
	}
	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) GetAll(ctx context.Context) ([]users.Domain, error) {
	var data []Users
	result := rep.Conn.Find(&data)
	if result.Error != nil {
		return []users.Domain{}, result.Error
	}

	return ToListDomain(data), nil
}

func (rep *MysqlUserRepository) GetById(ctx context.Context, Id uint) (users.Domain, error) {
	var user Users
	result := rep.Conn.Find(&user, "id = ?", Id)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Update(ctx context.Context, domain users.Domain, id uint) (users.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return users.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *MysqlUserRepository) Delete(ctx context.Context, id uint) error {
	user := Users{}
	err := rep.Conn.Delete(&user, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
