package karyawans

import (
	"alterra/business/karyawans"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MysqlKaryawanRepository struct {
	Conn *gorm.DB
}

func NewMysqlKaryawanRepository(conn *gorm.DB) *MysqlKaryawanRepository {
	return &MysqlKaryawanRepository{
		Conn: conn,
	}
}

func (rep *MysqlKaryawanRepository) GetByEmail(ctx context.Context, email string) (karyawans.Domain, error) {
	var karyawan Karyawans
	err := rep.Conn.Find(&karyawan, "email = ?", email)
	if err.Error != nil {
		return karyawans.Domain{}, err.Error
	}
	return karyawan.ToDomain(), nil
}

func (rep *MysqlKaryawanRepository) Register(ctx context.Context, domain *karyawans.Domain) (karyawans.Domain, error) {
	user := Karyawans{
		Id:       domain.Id,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}
	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return karyawans.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlKaryawanRepository) GetAll(ctx context.Context) ([]karyawans.Domain, error) {
	var data []Karyawans
	result := rep.Conn.Find(&data)
	if result.Error != nil {
		return []karyawans.Domain{}, result.Error
	}

	return ToListDomain(data), nil
}

func (rep *MysqlKaryawanRepository) GetById(ctx context.Context, Id uint) (karyawans.Domain, error) {
	var karyawan Karyawans
	result := rep.Conn.Find(&karyawan, "id = ?", Id)
	if result.Error != nil {
		return karyawans.Domain{}, result.Error
	}

	return karyawan.ToDomain(), nil
}

func (rep *MysqlKaryawanRepository) Update(ctx context.Context, domain karyawans.Domain, id uint) (karyawans.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return karyawans.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *MysqlKaryawanRepository) Delete(ctx context.Context, id uint) error {
	karyawan := Karyawans{}
	err := rep.Conn.Delete(&karyawan, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
