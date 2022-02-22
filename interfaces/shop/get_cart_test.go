package shop

import (
	"backend-ft/common/shop"
	mock_shop_application "backend-ft/mocks/application/shop"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shopInterface_GetCartByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("error", func(t *testing.T) {
		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetCartByUserIDAndStatusInCart(int64(1)).Return(nil, errors.New("failed"))
		si := NewShopInterface(nil, sa)
		_, err := si.GetCartByUserID(int64(1))
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		sa := mock_shop_application.NewMockShopApplication(ctrl)
		sa.EXPECT().GetCartByUserIDAndStatusInCart(int64(1)).Return([]shop.Item{}, nil)
		si := NewShopInterface(nil, sa)
		_, err := si.GetCartByUserID(int64(1))
		assert.Nil(t, err)
	})
}
