package shop

import (
	mock_shop_repository "backend-ft/mocks/repository/shop"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShopApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sr := mock_shop_repository.NewMockShopRepository(ctrl)
	sa := NewShopApplication(nil, sr)
	assert.IsType(t, &shopApplication{}, sa)
}
