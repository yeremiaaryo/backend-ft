package shop

import (
	"backend-ft/common/shop"
	mock_shop_repository "backend-ft/mocks/repository/shop"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shopApplication_InsertCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var (
		newDB  = sqlx.NewDb(db, "sqlMock")
		userID = int64(1)
		cart   = shop.Cart{}
	)

	t.Run("success", func(t *testing.T) {
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().InsertCart(userID, cart, nil).Return(errors.New("failed"))

		sa := NewShopApplication(newDB, sr)
		err := sa.InsertCart(userID, cart, nil)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().InsertCart(userID, cart, nil).Return(nil)

		sa := NewShopApplication(newDB, sr)
		err := sa.InsertCart(userID, cart, nil)
		assert.Nil(t, err)
	})
}
