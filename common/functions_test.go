package common

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestSystemResponse(t *testing.T) {
	t.Run("when err nil", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		assert.Nil(t, SystemResponse(ctx, "", nil, 200))
	})

	t.Run("when err bad request", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		assert.Nil(t, SystemResponse(ctx, "", errors.New("failed"), 400))
	})

	t.Run("when err 500", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		assert.Nil(t, SystemResponse(ctx, "", errors.New("failed"), 500))
	})
}
