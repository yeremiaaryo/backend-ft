package shop

import (
	shopDto "backend-ft/common/shop"
	"backend-ft/repository/shop"
	"github.com/jmoiron/sqlx"
)

type ShopApplication interface {
	GetItemByMultipleSKU(sku []string, tx *sqlx.Tx) ([]shopDto.Item, error)
	GetItemBySKU(sku string) (*shopDto.Item, error)
	InsertCheckout(userId int64, items []shopDto.Item, tx *sqlx.Tx) error
	UpdateItemQty(items []shopDto.Item, tx *sqlx.Tx) error
	GetCartByUserIDAndStatusInCart(userID int64) ([]shopDto.Item, error)
	GetCartByUserIDAndSKU(userID int64, cart shopDto.Cart, tx *sqlx.Tx) (*shopDto.Cart, error)
	InsertCart(userId int64, cart shopDto.Cart, tx *sqlx.Tx) error
	UpdateCart(userId int64, cart shopDto.Cart, status shopDto.CartStatus, tx *sqlx.Tx) error
	UpdateCartStatus(userId int64, sku string, status shopDto.CartStatus, tx *sqlx.Tx) error
}

type shopApplication struct {
	shopRepository shop.ShopRepository
	db             *sqlx.DB
}

func NewShopApplication(db *sqlx.DB, sr shop.ShopRepository) ShopApplication {
	return &shopApplication{
		shopRepository: sr,
		db:             db,
	}
}
