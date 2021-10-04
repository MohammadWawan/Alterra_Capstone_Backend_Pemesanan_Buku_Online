package users

import (
	"alterra/business/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Address   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		Name:      user.Name,
		Address:   user.Address,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:        domain.Id,
		Name:      domain.Name,
		Address:   domain.Address,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(data []Users) []users.Domain {
	list := []users.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
