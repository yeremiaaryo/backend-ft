package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
)

type ShopRepository interface {
	GetItemByMultipleSKU(sku []string, tx *sqlx.Tx) ([]shop.Item, error)
	GetItemBySKU(sku string) (*shop.Item, error)
	InsertCheckout(userId int64, item shop.Item, tx *sqlx.Tx) error
	UpdateItemQty(items []shop.Item, tx *sqlx.Tx) error
	GetCartByUserIDAndStatusInCart(userID int64) ([]shop.Cart, error)
	GetCartByUserIDAndSKU(userID int64, cart shop.Cart, tx *sqlx.Tx) (*shop.Cart, error)
	InsertCart(userId int64, cart shop.Cart, tx *sqlx.Tx) error
	UpdateCart(userId int64, cart shop.Cart, status shop.CartStatus, tx *sqlx.Tx) error
	UpdateCartStatus(userId int64, sku string, status shop.CartStatus, tx *sqlx.Tx) error
}

type shopRepository struct {
	db *sqlx.DB
}

func NewShopRepository(db *sqlx.DB) ShopRepository {
	return &shopRepository{
		db: db,
	}
}
