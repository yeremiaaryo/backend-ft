package shop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShopController(t *testing.T) {
	uc := NewShopController(nil)
	assert.IsType(t, &ShopController{}, uc)
}
