package transaction_details

import (
	"alterra/business/transaction_details"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MysqlTransaction_DetailRepository struct {
	Conn *gorm.DB
}

func NewMysqlTransaction_DetailRepository(conn *gorm.DB) *MysqlTransaction_DetailRepository {
	return &MysqlTransaction_DetailRepository{
		Conn: conn,
	}
}

func (rep *MysqlTransaction_DetailRepository) GetListTransaction_Details(ctx context.Context, search string) ([]transaction_details.Domain, error) {
	var data []Transaction_Detail
	err := rep.Conn.Find(&data)
	if err.Error != nil {
		return []transaction_details.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *MysqlTransaction_DetailRepository) GetById(ctx context.Context, Id uint) (transaction_details.Domain, error) {
	var transaction_detail Transaction_Detail
	result := rep.Conn.Find(&transaction_detail, "id = ?", Id)
	if result.Error != nil {
		return transaction_details.Domain{}, result.Error
	}

	return transaction_detail.ToDomain(), nil
}

func (rep *MysqlTransaction_DetailRepository) InsertTransaction_Details(ctx context.Context, domain transaction_details.Domain) (transaction_details.Domain, error) {
	transaction_detail := FromDomain(domain)
	err := rep.Conn.Create(&transaction_detail)
	if err.Error != nil {
		return transaction_details.Domain{}, err.Error
	}
	return transaction_detail.ToDomain(), nil
}

func (rep *MysqlTransaction_DetailRepository) Update(ctx context.Context, domain transaction_details.Domain, id uint) (transaction_details.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return transaction_details.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *MysqlTransaction_DetailRepository) Delete(ctx context.Context, id uint) error {
	transaction_detail := Transaction_Detail{}
	err := rep.Conn.Delete(&transaction_detail, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
