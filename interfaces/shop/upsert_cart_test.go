package shop

import (
	"backend-ft/common/shop"
	mock_shop_application "backend-ft/mocks/application/shop"
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func Test_shopInterface_UpsertCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("error no item", func(t *testing.T) {
		db, mock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin().WillReturnError(errors.New("failed"))
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.Cart{}
		userID := int64(1)

		sa := mock_shop_application.NewMockShopApplication(ctrl)

		si := NewShopInterface(db, sa)
		err = si.UpsertCart(userID, request)
		assert.Error(t, err)
	})

	t.Run("error no item", func(t *testing.T) {
		db, mock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.Cart{}
		userID := int64(1)

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemBySKU(request.SKU).Return(nil, sql.ErrNoRows)

		si := NewShopInterface(db, sa)
		err = si.UpsertCart(userID, request)
		assert.Error(t, err)
	})

	t.Run("error GetItemBySKU", func(t *testing.T) {
		db, mock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.Cart{}
		userID := int64(1)

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemBySKU(request.SKU).Return(nil, errors.New("failed"))

		si := NewShopInterface(db, sa)
		err = si.UpsertCart(userID, request)
		assert.Error(t, err)
	})

	t.Run("error item sold out", func(t *testing.T) {
		db, mock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.Cart{}
		userID := int64(1)

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemBySKU(request.SKU).Return(&shop.Item{Qty: 0}, nil)

		si := NewShopInterface(db, sa)
		err = si.UpsertCart(userID, request)
		assert.Error(t, err)
	})

	t.Run("error GetCartByUserIDAndSKU", func(t *testing.T) {
		db, mock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.Cart{}
		userID := int64(1)

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemBySKU(request.SKU).Return(&shop.Item{Qty: 1}, nil)
		sa.EXPECT().GetCartByUserIDAndSKU(userID, request, gomock.Any()).Return(nil, errors.New("failed"))

		si := NewShopInterface(db, sa)
		err = si.UpsertCart(userID, request)
		assert.Error(t, err)
	})

	t.Run("error insert", func(t *testing.T) {
		db, mock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.Cart{}
		userID := int64(1)

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemBySKU(request.SKU).Return(&shop.Item{Qty: 1}, nil)
		sa.EXPECT().GetCartByUserIDAndSKU(userID, request, gomock.Any()).Return(nil, sql.ErrNoRows)
		sa.EXPECT().InsertCart(userID, request, gomock.Any()).Return(errors.New("failed"))

		si := NewShopInterface(db, sa)
		err = si.UpsertCart(userID, request)
		assert.Error(t, err)
	})

	t.Run("error update", func(t *testing.T) {
		db, mock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.Cart{}
		userID := int64(1)

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemBySKU(request.SKU).Return(&shop.Item{Qty: 1}, nil)
		sa.EXPECT().GetCartByUserIDAndSKU(userID, request, gomock.Any()).Return(nil, nil)
		sa.EXPECT().UpdateCart(userID, request, shop.CartStatusInCart, gomock.Any()).Return(errors.New("failed"))

		si := NewShopInterface(db, sa)
		err = si.UpsertCart(userID, request)
		assert.Error(t, err)
	})

	t.Run("success insert", func(t *testing.T) {
		db, mock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.Cart{}
		userID := int64(1)

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemBySKU(request.SKU).Return(&shop.Item{Qty: 1}, nil)
		sa.EXPECT().GetCartByUserIDAndSKU(userID, request, gomock.Any()).Return(nil, sql.ErrNoRows)
		sa.EXPECT().InsertCart(userID, request, gomock.Any()).Return(nil)

		si := NewShopInterface(db, sa)
		err = si.UpsertCart(userID, request)
		assert.Nil(t, err)
	})

	t.Run("success update", func(t *testing.T) {
		db, mock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		mock.ExpectBegin()
		mock.ExpectCommit()
		mock.ExpectRollback()
		request := shop.Cart{}
		userID := int64(1)

		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetItemBySKU(request.SKU).Return(&shop.Item{Qty: 1}, nil)
		sa.EXPECT().GetCartByUserIDAndSKU(userID, request, gomock.Any()).Return(nil, nil)
		sa.EXPECT().UpdateCart(userID, request, shop.CartStatusInCart, gomock.Any()).Return(nil)

		si := NewShopInterface(db, sa)
		err = si.UpsertCart(userID, request)
		assert.Nil(t, err)
	})
}
