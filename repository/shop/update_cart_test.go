package shop

import (
	"backend-ft/common/shop"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shopRepository_UpdateCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("error prepare", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			userID = int64(1)
			newDB  = sqlx.NewDb(db, "sqlMock")
			cart   = shop.Cart{}
		)

		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		mock.ExpectPrepare(`UPDATE cart SET qty = ?, status = ? WHERE user_id = ? AND sku = ?`).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		err = sr.UpdateCart(userID, cart, shop.CartStatusInCart, tx)
		assert.Error(t, err)
	})

	t.Run("error exec", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			userID = int64(1)
			newDB  = sqlx.NewDb(db, "sqlMock")
			cart   = shop.Cart{}
		)
		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		stmt := mock.ExpectPrepare(`UPDATE cart SET qty = ?, status = ? WHERE user_id = ? AND sku = ?`)

		stmt.ExpectExec().WithArgs(cart.Qty, shop.CartStatusInCart, userID, cart.SKU).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		err = sr.UpdateCart(userID, cart, shop.CartStatusInCart, tx)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			userID = int64(1)
			newDB  = sqlx.NewDb(db, "sqlMock")
			cart   = shop.Cart{}
		)
		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		stmt := mock.ExpectPrepare(`UPDATE cart SET qty = ?, status = ? WHERE user_id = ? AND sku = ?`)

		stmt.ExpectExec().WithArgs(cart.Qty, shop.CartStatusInCart, userID, cart.SKU).WillReturnResult(sqlmock.NewResult(0, 1))
		sr := NewShopRepository(newDB)
		err = sr.UpdateCart(userID, cart, shop.CartStatusInCart, tx)
		assert.Nil(t, err)
	})
}

func Test_shopRepository_UpdateCartStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("error prepare", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			userID = int64(1)
			newDB  = sqlx.NewDb(db, "sqlMock")
			sku    = "sku"
		)

		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		mock.ExpectPrepare(`UPDATE cart SET status = ? WHERE user_id = ? AND sku = ?`).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		err = sr.UpdateCartStatus(userID, sku, shop.CartStatusInCart, tx)
		assert.Error(t, err)
	})

	t.Run("error exec", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			userID = int64(1)
			newDB  = sqlx.NewDb(db, "sqlMock")
			sku    = "sku"
		)
		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		stmt := mock.ExpectPrepare(`UPDATE cart SET status = ? WHERE user_id = ? AND sku = ?`)

		stmt.ExpectExec().WithArgs(shop.CartStatusInCart, userID, sku).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		err = sr.UpdateCartStatus(userID, sku, shop.CartStatusInCart, tx)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			userID = int64(1)
			newDB  = sqlx.NewDb(db, "sqlMock")
			sku    = "sku"
		)
		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		stmt := mock.ExpectPrepare(`UPDATE cart SET status = ? WHERE user_id = ? AND sku = ?`)

		stmt.ExpectExec().WithArgs(shop.CartStatusInCart, userID, sku).WillReturnResult(sqlmock.NewResult(0, 1))
		sr := NewShopRepository(newDB)
		err = sr.UpdateCartStatus(userID, sku, shop.CartStatusInCart, tx)
		assert.Nil(t, err)
	})
}
