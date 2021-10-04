package karyawans

import (
	"alterra/business/karyawans"
	"time"

	"gorm.io/gorm"
)

type Karyawans struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (karyawan *Karyawans) ToDomain() karyawans.Domain {
	return karyawans.Domain{
		Id:        karyawan.Id,
		Name:      karyawan.Name,
		Email:     karyawan.Email,
		Password:  karyawan.Password,
		CreatedAt: karyawan.CreatedAt,
		UpdatedAt: karyawan.UpdatedAt,
	}
}

func FromDomain(domain karyawans.Domain) Karyawans {
	return Karyawans{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToListDomain(data []Karyawans) []karyawans.Domain {
	list := []karyawans.Domain{}
	for _, v := range data {
		list = append(list, v.ToDomain())
	}

	return list
}
