package shop

import (
	"backend-ft/common/shop"
	"github.com/jmoiron/sqlx"
)

func (sa *shopApplication) UpdateItemQty(items []shop.Item, tx *sqlx.Tx) error {
	return sa.shopRepository.UpdateItemQty(items, tx)
}
