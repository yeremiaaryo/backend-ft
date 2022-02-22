package shop

import (
	"backend-ft/common/shop"
	"errors"
	"fmt"
)

func (si *shopInterface) Checkout(input shop.CheckoutRequest) (*shop.CheckoutResponse, error) {
	if len(input.Items) < 1 {
		return nil, errors.New("no items on checkout")
	}
	tx, err := si.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	sku := make([]string, 0)
	for key, value := range input.Items {
		if value == 0 {
			return nil, errors.New("invalid qty on items")
		}
		sku = append(sku, key)
	}
	items, err := si.shopApplication.GetItemByMultipleSKU(sku, tx)
	if err != nil {
		return nil, err
	}

	sum := float32(0)
	mapPrice := make(map[string]float32, 0)
	for _, v := range items {
		mapPrice[v.SKU] = v.Price
		if v.Qty < input.Items[v.SKU] {
			return nil, errors.New(fmt.Sprintf("Item %s doesn't have enough stock", v.SKU))
		}
		sum = sum + (float32(input.Items[v.SKU]) * v.Price)
		sum = sum * 100 / 100
	}

	discount := countDiscount(input, mapPrice)
	totalPrice := (sum - discount) * 100 / 100

	checkoutData := make([]shop.Item, 0)
	updateQtyData := make([]shop.Item, 0)
	for _, v := range items {
		val := input.Items[v.SKU]
		if val > 0 {
			item := shop.Item{
				SKU:   v.SKU,
				Name:  v.Name,
				Price: v.Price,
				Qty:   val,
			}
			checkoutData = append(checkoutData, item)

			updateQty := shop.Item{
				SKU:   v.SKU,
				Name:  v.Name,
				Price: v.Price,
				Qty:   v.Qty - val,
			}
			updateQtyData = append(updateQtyData, updateQty)

			err = si.shopApplication.UpdateCartStatus(input.UserID, v.SKU, shop.CartStatusDone, tx)
			if err != nil {
				return nil, err
			}
		}
	}

	err = si.shopApplication.InsertCheckout(input.UserID, checkoutData, tx)
	if err != nil {
		return nil, err
	}

	err = si.shopApplication.UpdateItemQty(updateQtyData, tx)
	if err != nil {
		return nil, err
	}

	return &shop.CheckoutResponse{
		Items:      checkoutData,
		TotalPrice: totalPrice,
	}, tx.Commit()
}

func countDiscount(input shop.CheckoutRequest, mapPrice map[string]float32) float32 {
	googleHomeDiscountQty := input.Items["120P90"] / 3
	googleHomeDiscount := mapPrice["120P90"] * float32(googleHomeDiscountQty)

	alexaDiscount := float32(0)
	if input.Items["A304SD"] >= 3 {
		alexaDiscount = float32(input.Items["A304SD"]) * mapPrice["A304SD"] * 0.1
	}

	raspberryDiscount := float32(0)
	if input.Items["43N23P"] >= input.Items["234234"] {
		raspberryDiscount = raspberryDiscount + (float32(input.Items["234234"]) * mapPrice["234234"])
	} else {
		raspberryDiscount = raspberryDiscount + (float32(input.Items["43N23P"]) * mapPrice["234234"])
	}
	return (googleHomeDiscount + alexaDiscount + raspberryDiscount) * 100 / 100
}
