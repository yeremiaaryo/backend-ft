package shop

import (
	"backend-ft/common/shop"
	"database/sql"
	"errors"
)

func (si *shopInterface) UpsertCart(userId int64, cart shop.Cart) error {
	tx, err := si.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	item, err := si.shopApplication.GetItemBySKU(cart.SKU)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("sku not available")
		}
		return err
	}

	if item != nil && item.Qty < 1 {
		return errors.New("item is sold out")
	}

	_, err = si.shopApplication.GetCartByUserIDAndSKU(userId, cart, tx)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		err = si.shopApplication.InsertCart(userId, cart, tx)
		if err != nil {
			return err
		}
	} else {
		err = si.shopApplication.UpdateCart(userId, cart, shop.CartStatusInCart, tx)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
