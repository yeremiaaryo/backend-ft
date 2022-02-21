package common

import (
	"github.com/labstack/echo"
	"net/http"
)

func SystemResponse(ec echo.Context, data interface{}, err error, code int) error {
	if err == nil {
		return ec.JSON(http.StatusOK, &BaseResponseDto{
			Code: http.StatusOK,
			Data: data,
		})
	}
	if code == http.StatusBadRequest {
		return ec.JSON(http.StatusBadRequest, &BaseResponseDto{
			Code:  http.StatusBadRequest,
			Data:  data,
			Error: err.Error(),
		})
	}
	return ec.JSON(http.StatusInternalServerError, &BaseResponseDto{
		Code:  http.StatusInternalServerError,
		Data:  data,
		Error: err.Error(),
	})
}
