package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
)

func (sa *shopApplication) UpdateCart(userId int64, cart shop.Cart, status shop.CartStatus, tx *sqlx.Tx) error {
	return sa.shopRepository.UpdateCart(userId, cart, status, tx)
}

func (sa *shopApplication) UpdateCartStatus(userId int64, sku string, status shop.CartStatus, tx *sqlx.Tx) error {
	return sa.shopRepository.UpdateCartStatus(userId, sku, status, tx)
}
