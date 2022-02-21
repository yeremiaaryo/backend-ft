package shop

import (
	"backend-ft/common"
	"backend-ft/common/shop"
	"github.com/labstack/echo"
	"net/http"
)

func (sc *ShopController) Checkout(ec echo.Context) error {
	var (
		request shop.CheckoutRequest
	)

	if err := ec.Bind(&request); err != nil {
		return common.SystemResponse(ec, nil, err, http.StatusBadRequest)
	}

	if err := ec.Validate(request); err != nil {
		return common.SystemResponse(ec, "invalid request", err, http.StatusBadRequest)
	}

	res, err := sc.shopInterface.Checkout(request)
	if err != nil {
		return common.SystemResponse(ec, nil, err, http.StatusBadRequest)
	}
	return common.SystemResponse(ec, res, nil, http.StatusOK)
}
