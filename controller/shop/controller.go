package shop

import (
	"backend-ft/interfaces/shop"
)

type ShopController struct {
	shopInterface shop.ShopInterface
}

func NewShopController(si shop.ShopInterface) *ShopController {
	return &ShopController{
		shopInterface: si,
	}
}
