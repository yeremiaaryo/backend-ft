package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
	"log"
)

func (sa *shopApplication) InsertCheckout(userId int64, items []shop.Item, tx *sqlx.Tx) error {
	for _, v := range items {
		err := sa.shopRepository.InsertCheckout(userId, v, tx)
		if err != nil {
			log.Println("error when exec statement, error:", err.Error())
			return err
		}
	}
	return nil
}
