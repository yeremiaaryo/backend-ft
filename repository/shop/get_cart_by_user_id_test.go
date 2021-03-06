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

func Test_shopRepository_GetCartByUserIDAndStatusInCart(t *testing.T) {
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
		)

		mock.ExpectPrepare(`SELECT sku, qty FROM cart WHERE user_id = ? AND status = ?`).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		_, err = sr.GetCartByUserIDAndStatusInCart(userID)
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
		)

		stmt := mock.ExpectPrepare(`SELECT sku, qty FROM cart WHERE user_id = ? AND status = ?`)
		rows := sqlmock.NewRows([]string{"sku", "qty"}).
			AddRow("sku", 1)
		stmt.ExpectQuery().WithArgs(userID, shop.CartStatusInCart).WillReturnRows(rows)
		sr := NewShopRepository(newDB)
		_, err = sr.GetCartByUserIDAndStatusInCart(userID)
		assert.Nil(t, err)
	})
}

func Test_shopRepository_GetCartByUserIDAndSKU(t *testing.T) {
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
		mock.ExpectPrepare(`SELECT sku, qty FROM cart WHERE user_id = ? AND sku = ? FOR UPDATE`).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		_, err = sr.GetCartByUserIDAndSKU(userID, cart, tx)
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
		stmt := mock.ExpectPrepare(`SELECT sku, qty FROM cart WHERE user_id = ? AND sku = ? FOR UPDATE`)
		rows := sqlmock.NewRows([]string{"sku", "qty"}).
			AddRow("sku", 1)
		stmt.ExpectQuery().WithArgs(userID, cart.SKU).WillReturnRows(rows)
		sr := NewShopRepository(newDB)
		_, err = sr.GetCartByUserIDAndSKU(userID, cart, tx)
		assert.Nil(t, err)
	})
}
