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

func TestShopController_UpsertCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("When error bind", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", ioutil.NopCloser(nil))
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		si := mock_shop_interface.NewMockShopInterface(ctrl)
		h := NewShopController(si)
		err := h.UpsertCart(ctx)
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
		err := h.UpsertCart(ctx)
		assert.Nil(t, err)
	})

	t.Run("When error upsert cart", func(t *testing.T) {
		e := echo.New()
		body := `{"sku":"sku","qty":1,"user_id":1}`
		req := httptest.NewRequest(echo.POST, "/", ioutil.NopCloser(bytes.NewBufferString(body)))
		req.Header.Add("Content-Type", "application/json")
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		request := shop.UpsertCartRequest{
			SKU:    "sku",
			Qty:    1,
			UserID: int64(1),
		}

		si := mock_shop_interface.NewMockShopInterface(ctrl)
		cart := shop.Cart{
			SKU: request.SKU,
			Qty: request.Qty,
		}
		si.EXPECT().UpsertCart(request.UserID, cart).Return(errors.New("failed"))
		h := NewShopController(si)
		e.Validator = &common.CustomValidator{
			Validator: validator.New(),
		}
		err := h.UpsertCart(ctx)
		assert.Nil(t, err)
	})

	t.Run("When success", func(t *testing.T) {
		e := echo.New()
		body := `{"sku":"sku","qty":1,"user_id":1}`
		req := httptest.NewRequest(echo.POST, "/", ioutil.NopCloser(bytes.NewBufferString(body)))
		req.Header.Add("Content-Type", "application/json")
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		request := shop.UpsertCartRequest{
			SKU:    "sku",
			Qty:    1,
			UserID: int64(1),
		}
		si := mock_shop_interface.NewMockShopInterface(ctrl)
		cart := shop.Cart{
			SKU: request.SKU,
			Qty: request.Qty,
		}
		si.EXPECT().UpsertCart(request.UserID, cart).Return(nil)
		h := NewShopController(si)
		e.Validator = &common.CustomValidator{
			Validator: validator.New(),
		}
		err := h.UpsertCart(ctx)
		assert.Nil(t, err)
	})
}

func TestShopController_GetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("When error invalid user id", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/?user_id=abc", nil)
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		si := mock_shop_interface.NewMockShopInterface(ctrl)
		h := NewShopController(si)
		e.Validator = &common.CustomValidator{
			Validator: validator.New(),
		}
		err := h.GetCart(ctx)
		assert.Nil(t, err)
	})

	t.Run("When error GetCartByUserID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/?user_id=1", nil)
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		si := mock_shop_interface.NewMockShopInterface(ctrl)
		si.EXPECT().GetCartByUserID(int64(1)).Return(nil, errors.New("failed"))
		h := NewShopController(si)
		e.Validator = &common.CustomValidator{
			Validator: validator.New(),
		}
		err := h.GetCart(ctx)
		assert.Nil(t, err)
	})

	t.Run("When success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/?user_id=1", nil)
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		si := mock_shop_interface.NewMockShopInterface(ctrl)
		si.EXPECT().GetCartByUserID(int64(1)).Return([]shop.Item{}, nil)
		h := NewShopController(si)
		e.Validator = &common.CustomValidator{
			Validator: validator.New(),
		}
		err := h.GetCart(ctx)
		assert.Nil(t, err)
	})
}
