package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
	"log"
)

func (sr *shopRepository) GetItemByMultipleSKU(sku []string, tx *sqlx.Tx) ([]shop.Item, error) {

	q, args, err := sqlx.In(getItemByMultipleSKU, sku)
	if err != nil {
		return nil, err
	}
	statement, err := tx.Preparex(q)
	if err != nil {
		log.Println("error exec statement, error:", err.Error())
		return nil, err
	}

	data := make([]shop.Item, 0)
	err = statement.Select(&data, args...)
	return data, err
}

func (sr *shopRepository) GetItemBySKU(sku string) (*shop.Item, error) {

	statement, err := sr.db.Preparex(getItemBySKU)
	if err != nil {
		log.Println("error exec statement, error:", err.Error())
		return nil, err
	}

	item := shop.Item{}
	err = statement.QueryRowx(sku).StructScan(&item)
	return &item, err
}
