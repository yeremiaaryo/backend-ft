package shop

import (
	"backend-ft/common/shop"
	mock_shop_repository "backend-ft/mocks/repository/shop"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shopApplication_GetCartByUserIDAndSKU(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var (
		userID = int64(1)
		cart   = shop.Cart{}
		newDB  = sqlx.NewDb(db, "sqlMock")
	)
	mock.ExpectBegin()
	mock.ExpectCommit()
	tx, err := newDB.Beginx()

	t.Run("error", func(t *testing.T) {
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetCartByUserIDAndSKU(userID, cart, tx).Return(nil, errors.New("failed"))

		sa := shopApplication{
			shopRepository: sr,
		}
		_, err := sa.GetCartByUserIDAndSKU(userID, cart, tx)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		resp := &shop.Cart{}
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetCartByUserIDAndSKU(userID, cart, tx).Return(resp, nil)

		sa := shopApplication{
			shopRepository: sr,
		}
		res, err := sa.GetCartByUserIDAndSKU(userID, cart, tx)
		expected := &shop.Cart{}
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})
}

func Test_shopApplication_GetCartByUserIDAndStatusInCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var userID = int64(1)
	t.Run("error", func(t *testing.T) {
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetCartByUserIDAndStatusInCart(userID).Return(nil, errors.New("failed"))

		sa := shopApplication{
			shopRepository: sr,
		}
		_, err := sa.GetCartByUserIDAndStatusInCart(userID)
		assert.Error(t, err)
	})

	t.Run("error no row", func(t *testing.T) {
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetCartByUserIDAndStatusInCart(userID).Return(nil, sql.ErrNoRows)

		sa := shopApplication{
			shopRepository: sr,
		}
		_, err := sa.GetCartByUserIDAndStatusInCart(userID)
		assert.Nil(t, err)
	})

	t.Run("success, error get sku", func(t *testing.T) {
		resp := []shop.Cart{
			{},
		}
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetCartByUserIDAndStatusInCart(userID).Return(resp, nil)
		sr.EXPECT().GetItemBySKU("").Return(nil, errors.New("failed"))

		sa := shopApplication{
			shopRepository: sr,
		}
		res, err := sa.GetCartByUserIDAndStatusInCart(userID)
		expected := make([]shop.Item, 0)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("success, error get sku", func(t *testing.T) {
		resp := []shop.Cart{
			{},
		}
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetCartByUserIDAndStatusInCart(userID).Return(resp, nil)
		sr.EXPECT().GetItemBySKU("").Return(&shop.Item{}, nil)

		sa := shopApplication{
			shopRepository: sr,
		}
		res, err := sa.GetCartByUserIDAndStatusInCart(userID)
		expected := []shop.Item{
			{},
		}
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

}
