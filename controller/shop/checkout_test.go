package shop

import (
	"backend-ft/common"
	"backend-ft/common/shop"
	mock_shop_interface "backend-ft/mocks/interfaces/shop"
	"bytes"
	"errors"
	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestShopController_Checkout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("When error bind", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", ioutil.NopCloser(nil))
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		si := mock_shop_interface.NewMockShopInterface(ctrl)
		h := NewShopController(si)
		err := h.Checkout(ctx)
		assert.Nil(t, err)
	})

	t.Run("When error no validator", func(t *testing.T) {
		e := echo.New()
		body := `{"test":"test"}`
		req := httptest.NewRequest(echo.POST, "/", ioutil.NopCloser(bytes.NewBufferString(body)))
		req.Header.Add("Content-Type", "application/json")
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		si := mock_shop_interface.NewMockShopInterface(ctrl)
		h := NewShopController(si)
		err := h.Checkout(ctx)
		assert.Nil(t, err)
	})

	t.Run("When error checkout", func(t *testing.T) {
		e := echo.New()
		body := `{"items":{"sku":1},"user_id":1}`
		req := httptest.NewRequest(echo.POST, "/", ioutil.NopCloser(bytes.NewBufferString(body)))
		req.Header.Add("Content-Type", "application/json")
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 1,
			},
			UserID: 1,
		}

		si := mock_shop_interface.NewMockShopInterface(ctrl)
		si.EXPECT().Checkout(request).Return(nil, errors.New("failed"))
		h := NewShopController(si)
		e.Validator = &common.CustomValidator{
			Validator: validator.New(),
		}
		err := h.Checkout(ctx)
		assert.Nil(t, err)
	})

	t.Run("When success", func(t *testing.T) {
		e := echo.New()
		body := `{"items":{"sku":1},"user_id":1}`
		req := httptest.NewRequest(echo.POST, "/", ioutil.NopCloser(bytes.NewBufferString(body)))
		req.Header.Add("Content-Type", "application/json")
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		request := shop.CheckoutRequest{
			Items: map[string]int{
				"sku": 1,
			},
			UserID: 1,
		}
		si := mock_shop_interface.NewMockShopInterface(ctrl)
		si.EXPECT().Checkout(request).Return(&shop.CheckoutResponse{}, nil)
		h := NewShopController(si)
		e.Validator = &common.CustomValidator{
			Validator: validator.New(),
		}
		err := h.Checkout(ctx)
		assert.Nil(t, err)
	})
}
