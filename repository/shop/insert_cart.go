package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
	"log"
)

func (sr *shopRepository) InsertCart(userId int64, cart shop.Cart, tx *sqlx.Tx) error {
	statement, err := tx.Preparex(insertCart)
	if err != nil {
		log.Println("error prepare statement, error:", err.Error())
		return err
	}

	_, err = statement.Exec(userId, cart.SKU, cart.Qty, shop.CartStatusInCart)
	if err != nil {
		log.Println("error when exec statement, error:", err.Error())
		return err
	}
	return nil
}
