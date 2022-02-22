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

func Test_shopRepository_InsertCheckout(t *testing.T) {
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
			item   = shop.Item{}
		)

		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		mock.ExpectPrepare(`INSERT INTO checkout (user_id, sku, name, price, qty) VALUES (?, ?, ?, ?, ?)`).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		err = sr.InsertCheckout(userID, item, tx)
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
			item   = shop.Item{}
		)
		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		stmt := mock.ExpectPrepare(`INSERT INTO checkout (user_id, sku, name, price, qty) VALUES (?, ?, ?, ?, ?)`)

		stmt.ExpectExec().WithArgs(userID, item.SKU, item.Name, item.Price, item.Qty).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		err = sr.InsertCheckout(userID, item, tx)
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
			item   = shop.Item{}
		)
		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		stmt := mock.ExpectPrepare(`INSERT INTO checkout (user_id, sku, name, price, qty) VALUES (?, ?, ?, ?, ?)`)

		stmt.ExpectExec().WithArgs(userID, item.SKU, item.Name, item.Price, item.Qty).WillReturnResult(sqlmock.NewResult(0, 1))
		sr := NewShopRepository(newDB)
		err = sr.InsertCheckout(userID, item, tx)
		assert.Nil(t, err)
	})
}
