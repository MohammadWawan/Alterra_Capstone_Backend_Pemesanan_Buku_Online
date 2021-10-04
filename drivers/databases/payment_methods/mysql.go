package payment_methods

import (
	"alterra/business/payment_methods"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MysqlPayment_MethodRepository struct {
	Conn *gorm.DB
}

func NewMysqlPayment_MethodRepository(conn *gorm.DB) *MysqlPayment_MethodRepository {
	return &MysqlPayment_MethodRepository{
		Conn: conn,
	}
}

func (rep *MysqlPayment_MethodRepository) GetListPayment_Method(ctx context.Context, search string) ([]payment_methods.Domain, error) {
	var data []Payment_Method
	err := rep.Conn.Find(&data)
	if err.Error != nil {
		return []payment_methods.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *MysqlPayment_MethodRepository) GetById(ctx context.Context, Id uint) (payment_methods.Domain, error) {
	var payment_method Payment_Method
	result := rep.Conn.Find(&payment_method, "id = ?", Id)
	if result.Error != nil {
		return payment_methods.Domain{}, result.Error
	}

	return payment_method.ToDomain(), nil
}

func (rep *MysqlPayment_MethodRepository) InsertPayment_Method(ctx context.Context, domain payment_methods.Domain) (payment_methods.Domain, error) {
	payment_method := FromDomain(domain)
	err := rep.Conn.Create(&payment_method)
	if err.Error != nil {
		return payment_methods.Domain{}, err.Error
	}
	return payment_method.ToDomain(), nil
}

func (rep *MysqlPayment_MethodRepository) Update(ctx context.Context, domain payment_methods.Domain, id uint) (payment_methods.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return payment_methods.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *MysqlPayment_MethodRepository) Delete(ctx context.Context, id uint) error {
	payment_method := Payment_Method{}
	err := rep.Conn.Delete(&payment_method, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
