package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
	"log"
)

func (sr *shopRepository) GetCartByUserIDAndStatusInCart(userID int64) ([]shop.Cart, error) {

	statement, err := sr.db.Preparex(getCartByUserIDAndStatus)
	if err != nil {
		log.Println("error exec statement, error:", err.Error())
		return nil, err
	}

	data := make([]shop.Cart, 0)
	err = statement.Select(&data, userID, shop.CartStatusInCart)
	return data, err
}

func (sr *shopRepository) GetCartByUserIDAndSKU(userID int64, cart shop.Cart, tx *sqlx.Tx) (*shop.Cart, error) {

	statement, err := tx.Preparex(getCartByUserIDAndSKUForUpdate)
	if err != nil {
		log.Println("error exec statement, error:", err.Error())
		return nil, err
	}

	data := shop.Cart{}
	err = statement.QueryRowx(userID, cart.SKU).StructScan(&data)
	return &data, err
}
