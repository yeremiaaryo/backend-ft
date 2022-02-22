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

func Test_shopApplication_GetItemByMultipleSKU(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var (
		newDB = sqlx.NewDb(db, "sqlMock")
	)
	mock.ExpectBegin()
	mock.ExpectCommit()
	tx, err := newDB.Beginx()

	t.Run("failed", func(t *testing.T) {
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetItemByMultipleSKU([]string{"sku"}, tx).Return(nil, errors.New("failed"))
		sa := shopApplication{
			shopRepository: sr,
		}
		_, err := sa.GetItemByMultipleSKU([]string{"sku"}, tx)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetItemByMultipleSKU([]string{"sku"}, tx).Return([]shop.Item{{}}, nil)
		sa := shopApplication{
			shopRepository: sr,
		}
		res, err := sa.GetItemByMultipleSKU([]string{"sku"}, tx)
		expected := []shop.Item{{}}
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})
}

func Test_shopApplication_GetItemBySKU(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("failed", func(t *testing.T) {
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetItemBySKU("sku").Return(nil, errors.New("failed"))
		sa := shopApplication{
			shopRepository: sr,
		}
		_, err := sa.GetItemBySKU("sku")
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		sr := mock_shop_repository.NewMockShopRepository(ctrl)
		sr.EXPECT().GetItemBySKU("sku").Return(&shop.Item{}, nil)
		sa := shopApplication{
			shopRepository: sr,
		}
		res, err := sa.GetItemBySKU("sku")
		expected := &shop.Item{}
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})
}
