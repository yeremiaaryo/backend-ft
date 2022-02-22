package shop

import (
	"backend-ft/common/shop"
	mock_shop_application "backend-ft/mocks/application/shop"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func Test_shopInterface_Checkout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("no items", func(t *testing.T) {

		request := shop.CheckoutRequest{
			Items:  map[string]int{},
			UserID: 1,
		}

		si := NewShopInterface(nil, nil)
		_, err := si.Checkout(request)

		assert.Error(t, err)
	})

	t.Run("error begin", func(t *testing.T) {
		db, mock, err := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin().WillReturnError(errors.New("failed"))
		mock.ExpectRollback()
		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 1,
			},
			UserID: 1,
		}

		si := NewShopInterface(db, nil)
		_, err = si.Checkout(request)

		assert.Error(t, err)
	})

	t.Run("invalid qty", func(t *testing.T) {
		db, mock, err := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectRollback()
		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 0,
			},
			UserID: 1,
		}

		si := NewShopInterface(db, nil)
		_, err = si.Checkout(request)

		assert.Error(t, err)
	})

	t.Run("error GetItemByMultipleSKU", func(t *testing.T) {
		db, mock, err := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectRollback()
		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 1,
			},
			UserID: 1,
		}

		sku := []string{"sku"}

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemByMultipleSKU(sku, gomock.Any()).Return(nil, errors.New("failed"))

		si := NewShopInterface(db, sa)
		_, err = si.Checkout(request)

		assert.Error(t, err)
	})

	t.Run("not enough stock", func(t *testing.T) {
		db, mock, err := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectRollback()
		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 1,
			},
			UserID: 1,
		}

		sku := []string{"sku"}

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemByMultipleSKU(sku, gomock.Any()).Return([]shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   0,
			},
		}, nil)

		si := NewShopInterface(db, sa)
		_, err = si.Checkout(request)

		assert.Error(t, err)
	})

	t.Run("error UpdateCartStatus", func(t *testing.T) {
		db, mock, err := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectRollback()
		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 1,
			},
			UserID: 1,
		}

		sku := []string{"sku"}

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemByMultipleSKU(sku, gomock.Any()).Return([]shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   2,
			},
		}, nil)
		sa.EXPECT().UpdateCartStatus(request.UserID, "sku", shop.CartStatusDone, gomock.Any()).Return(errors.New("failed"))

		si := NewShopInterface(db, sa)
		_, err = si.Checkout(request)

		assert.Error(t, err)
	})

	t.Run("error InsertCheckout", func(t *testing.T) {
		db, mock, err := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectRollback()
		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 1,
			},
			UserID: 1,
		}

		sku := []string{"sku"}

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemByMultipleSKU(sku, gomock.Any()).Return([]shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   2,
			},
		}, nil)
		sa.EXPECT().UpdateCartStatus(request.UserID, "sku", shop.CartStatusDone, gomock.Any()).Return(nil)
		checkoutData := []shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   1,
			},
		}
		sa.EXPECT().InsertCheckout(request.UserID, checkoutData, gomock.Any()).Return(errors.New("failed"))

		si := NewShopInterface(db, sa)
		_, err = si.Checkout(request)

		assert.Error(t, err)
	})

	t.Run("error UpdateItemQty", func(t *testing.T) {
		db, mock, err := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectRollback()
		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 1,
			},
			UserID: 1,
		}

		sku := []string{"sku"}

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemByMultipleSKU(sku, gomock.Any()).Return([]shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   2,
			},
		}, nil)
		sa.EXPECT().UpdateCartStatus(request.UserID, "sku", shop.CartStatusDone, gomock.Any()).Return(nil)
		checkoutData := []shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   1,
			},
		}
		sa.EXPECT().InsertCheckout(request.UserID, checkoutData, gomock.Any()).Return(nil)
		updateData := []shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   1,
			},
		}
		sa.EXPECT().UpdateItemQty(updateData, gomock.Any()).Return(errors.New("failed"))

		si := NewShopInterface(db, sa)
		_, err = si.Checkout(request)

		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 1,
			},
			UserID: 1,
		}

		sku := []string{"sku"}

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemByMultipleSKU(sku, gomock.Any()).Return([]shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   2,
			},
		}, nil)
		sa.EXPECT().UpdateCartStatus(request.UserID, "sku", shop.CartStatusDone, gomock.Any()).Return(nil)
		checkoutData := []shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   1,
			},
		}
		sa.EXPECT().InsertCheckout(request.UserID, checkoutData, gomock.Any()).Return(nil)
		updateData := []shop.Item{
			{
				SKU:   "sku",
				Name:  "sku thing",
				Price: 1000,
				Qty:   1,
			},
		}
		sa.EXPECT().UpdateItemQty(updateData, gomock.Any()).Return(nil)

		si := NewShopInterface(db, sa)
		res, err := si.Checkout(request)
		expected := &shop.CheckoutResponse{
			Items: []shop.Item{
				{
					SKU:   "sku",
					Name:  "sku thing",
					Price: 1000,
					Qty:   1,
				},
			},
			TotalPrice: 1000,
		}
		assert.Equal(t, expected, res)
		assert.Nil(t, err)
	})
}

func Test_countDiscount(t *testing.T) {
	input := shop.CheckoutRequest{
		Items: map[string]int{
			"120P90": 3,
			"A304SD": 3,
			"234234": 1,
		},
	}
	mapPrice := map[string]float32{
		"120P90": 49.99,
		"A304SD": 109.50,
		"234234": 30,
	}
	res := countDiscount(input, mapPrice)
	assert.Equal(t, float32(82.84), res)
}
