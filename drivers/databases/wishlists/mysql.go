package wishlists

import (
	"alterra/business/wishlists"
	"context"
	"errors"

	"gorm.io/gorm"
)

type MysqlWishlistRepository struct {
	Conn *gorm.DB
}

func NewMysqlWishlistRepository(conn *gorm.DB) *MysqlWishlistRepository {
	return &MysqlWishlistRepository{
		Conn: conn,
	}
}

func (rep *MysqlWishlistRepository) GetListWishlist(ctx context.Context, search string) ([]wishlists.Domain, error) {
	var data []Wishlist
	err := rep.Conn.Find(&data)
	if err.Error != nil {
		return []wishlists.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *MysqlWishlistRepository) GetById(ctx context.Context, Id uint) (wishlists.Domain, error) {
	var wishlist Wishlist
	result := rep.Conn.Find(&wishlist, "id = ?", Id)
	if result.Error != nil {
		return wishlists.Domain{}, result.Error
	}

	return wishlist.ToDomain(), nil
}

func (rep *MysqlWishlistRepository) InsertWishlist(ctx context.Context, domain wishlists.Domain) (wishlists.Domain, error) {
	wishlist := FromDomain(domain)
	err := rep.Conn.Create(&wishlist)
	if err.Error != nil {
		return wishlists.Domain{}, err.Error
	}
	return wishlist.ToDomain(), nil
}

func (rep *MysqlWishlistRepository) Update(ctx context.Context, domain wishlists.Domain, id uint) (wishlists.Domain, error) {
	data := FromDomain(domain)
	if rep.Conn.Save(&data).Error != nil {
		return wishlists.Domain{}, errors.New("bad request")
	}

	return data.ToDomain(), nil
}

func (rep *MysqlWishlistRepository) Delete(ctx context.Context, id uint) error {
	wishlist := Wishlist{}
	err := rep.Conn.Delete(&wishlist, id)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}

	return nil
}
