package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
)

func (sa *shopApplication) GetItemByMultipleSKU(sku []string, tx *sqlx.Tx) ([]shop.Item, error) {
	return sa.shopRepository.GetItemByMultipleSKU(sku, tx)
}

func (sa *shopApplication) GetItemBySKU(sku string) (*shop.Item, error) {
	return sa.shopRepository.GetItemBySKU(sku)
}
