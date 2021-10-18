package transactions

import (
	"alterra/business/transactions"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMysqlTransactionRepository(conn *gorm.DB) *MysqlTransactionRepository {
	return &MysqlTransactionRepository{
		Conn: conn,
	}
}

func (rep *MysqlTransactionRepository) GetListTransactions(ctx context.Context, Method_Payment_Id uint, User_Id uint, Karyawan_Id uint) ([]transactions.Domain, error) {
	var data []Transaction
	err := rep.Conn.Find(&data)
	if err.Error != nil {
		return []transactions.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *MysqlTransactionRepository) GetById(ctx context.Context, Id uint) (transactions.Domain, error) {
	var transaction Transaction
	result := rep.Conn.Find(&transaction, "id = ?", Id)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	return transaction.ToDomain(), nil
}

func (rep *MysqlTransactionRepository) InsertTransactions(ctx context.Context, domain *transactions.Domain) (transactions.Domain, error) {
	transaction := FromDomain(*domain)
	err := rep.Conn.Create(&transaction)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return transaction.ToDomain(), nil
}

func (rep *MysqlTransactionRepository) Update(ctx context.Context, domain transactions.Domain, id uint) (transactions.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return transactions.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *MysqlTransactionRepository) Delete(ctx context.Context, id uint) error {
	transaction := Transaction{}
	err := rep.Conn.Delete(&transaction, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
