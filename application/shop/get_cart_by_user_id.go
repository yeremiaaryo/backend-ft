package shop

import (
	"backend-ft/common/shop"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func (sa *shopApplication) GetCartByUserIDAndStatusInCart(userID int64) ([]shop.Item, error) {
	cart, err := sa.shopRepository.GetCartByUserIDAndStatusInCart(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return sa.buildCartToItem(cart), err
}

func (sa *shopApplication) buildCartToItem(cart []shop.Cart) []shop.Item {
	items := make([]shop.Item, 0)
	for _, v := range cart {
		i, err := sa.shopRepository.GetItemBySKU(v.SKU)
		if err != nil {
			continue
		}
		if i != nil {
			item := shop.Item{
				SKU:   v.SKU,
				Name:  i.Name,
				Price: i.Price,
				Qty:   v.Qty,
			}
			items = append(items, item)
		}
	}
	return items
}

func (sa *shopApplication) GetCartByUserIDAndSKU(userID int64, cart shop.Cart, tx *sqlx.Tx) (*shop.Cart, error) {
	item, err := sa.shopRepository.GetCartByUserIDAndSKU(userID, cart, tx)
	if err != nil {
		return nil, err
	}

	return item, err
}
