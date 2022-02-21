package shop

import (
	mock_shop_application "backend-ft/mocks/application/shop"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShopInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sa := mock_shop_application.NewMockShopApplication(ctrl)
	interFace := NewShopInterface(nil, sa)
	assert.IsType(t, &shopInterface{}, interFace)
}
