package shop

import "backend-ft/common/shop"

func (si *shopInterface) GetCartByUserID(userID int64) ([]shop.Item, error) {
	return si.shopApplication.GetCartByUserIDAndStatusInCart(userID)
}
