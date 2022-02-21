package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
	"log"
)

func (sr *shopRepository) InsertCheckout(userId int64, item shop.Item, tx *sqlx.Tx) error {
	statement, err := tx.Preparex(insertCheckout)
	if err != nil {
		log.Println("error prepare statement, error:", err.Error())
		return err
	}

	_, err = statement.Exec(userId, item.SKU, item.Name, item.Price, item.Qty)
	if err != nil {
		log.Println("error when exec statement, error:", err.Error())
		return err
	}

	return err
}
