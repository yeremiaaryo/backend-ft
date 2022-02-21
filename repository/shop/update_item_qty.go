package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
	"log"
)

func (sr *shopRepository) UpdateItemQty(items []shop.Item, tx *sqlx.Tx) error {
	statement, err := tx.Preparex(updateItemQty)
	if err != nil {
		log.Println("error prepare statement, error:", err.Error())
		return err
	}

	for _, v := range items {
		_, err = statement.Exec(v.Qty, v.SKU)
		if err != nil {
			log.Println("error when exec statement, error:", err.Error())
			return err
		}
	}

	return err
}
