package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
)

func (sa *shopApplication) InsertCart(userId int64, cart shop.Cart, tx *sqlx.Tx) error {
	return sa.shopRepository.InsertCart(userId, cart, tx)
}
