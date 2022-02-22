package shop

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_shopRepository_GetCartByUserIDAndSKU(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

}

func Test_shopRepository_GetCartByUserIDAndStatusInCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

}
