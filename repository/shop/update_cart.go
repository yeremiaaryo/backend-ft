package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
	"log"
)

func (sr *shopRepository) UpdateCart(userId int64, cart shop.Cart, status shop.CartStatus, tx *sqlx.Tx) error {
	statement, err := tx.Preparex(updateCart)
	if err != nil {
		log.Println("error prepare statement, error:", err.Error())
		return err
	}

	_, err = statement.Exec(cart.Qty, status, userId, cart.SKU)
	if err != nil {
		log.Println("error when exec statement, error:", err.Error())
		return err
	}
	return nil
}

func (sr *shopRepository) UpdateCartStatus(userId int64, sku string, status shop.CartStatus, tx *sqlx.Tx) error {
	statement, err := tx.Preparex(updateCartStatus)
	if err != nil {
		log.Println("error prepare statement, error:", err.Error())
		return err
	}

	_, err = statement.Exec(status, userId, sku)
	if err != nil {
		log.Println("error when exec statement, error:", err.Error())
		return err
	}
	return nil
}
