package shop

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shopRepository_GetItemByMultipleSKU(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("error prepare", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			newDB = sqlx.NewDb(db, "sqlMock")
			sku   = []string{""}
		)

		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		q, _, _ := sqlx.In(`SELECT sku, name, price, qty FROM items WHERE sku IN (?) FOR UPDATE`, sku)
		mock.ExpectPrepare(q).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		_, err = sr.GetItemByMultipleSKU(sku, tx)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			newDB = sqlx.NewDb(db, "sqlMock")
			sku   = []string{""}
		)
		mock.ExpectBegin()
		tx, _ := newDB.Beginx()
		q, _, _ := sqlx.In(`SELECT sku, name, price, qty FROM items WHERE sku IN (?) FOR UPDATE`, sku)
		stmt := mock.ExpectPrepare(q)
		rows := sqlmock.NewRows([]string{"sku", "name", "price", "qty"}).
			AddRow("sku", "name", 29.99, 1)
		stmt.ExpectQuery().WithArgs("").WillReturnRows(rows)
		sr := NewShopRepository(newDB)
		_, err = sr.GetItemByMultipleSKU(sku, tx)
		assert.Nil(t, err)
	})
}

func Test_shopRepository_GetItemBySKU(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("error prepare", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			sku   = "sku"
			newDB = sqlx.NewDb(db, "sqlMock")
		)

		mock.ExpectPrepare(`SELECT sku, name, price, qty FROM items WHERE sku = ?`).WillReturnError(errors.New("failed"))
		sr := NewShopRepository(newDB)
		_, err = sr.GetItemBySKU(sku)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		var (
			sku   = "sku"
			newDB = sqlx.NewDb(db, "sqlMock")
		)

		stmt := mock.ExpectPrepare(`SELECT sku, name, price, qty FROM items WHERE sku = ?`)
		rows := sqlmock.NewRows([]string{"sku", "name", "price", "qty"}).
			AddRow("sku", "name", 29.99, 1)
		stmt.ExpectQuery().WithArgs(sku).WillReturnRows(rows)
		sr := NewShopRepository(newDB)
		_, err = sr.GetItemBySKU(sku)
		assert.Nil(t, err)
	})
}
