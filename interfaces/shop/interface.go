package shop

import (
	"backend-ft/application/shop"
	shopDto "backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
)

type ShopInterface interface {
	Checkout(input shopDto.CheckoutRequest) (*shopDto.CheckoutResponse, error)
	GetCartByUserID(userID int64) ([]shopDto.Item, error)
	UpsertCart(userId int64, cart shopDto.Cart) error
}

type shopInterface struct {
	shopApplication shop.ShopApplication
	db              *sqlx.DB
}

func NewShopInterface(db *sqlx.DB, sa shop.ShopApplication) ShopInterface {
	return &shopInterface{
		shopApplication: sa,
		db:              db,
	}
}
