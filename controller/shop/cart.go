package shop

import (
	"backend-ft/common"
	"backend-ft/common/shop"
	"errors"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (sc *ShopController) GetCart(ec echo.Context) error {

	userIDStr := ec.QueryParam("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return common.SystemResponse(ec, nil, errors.New("invalid user id"), http.StatusBadRequest)
	}
	res, err := sc.shopInterface.GetCartByUserID(userID)
	if err != nil {
		return common.SystemResponse(ec, nil, err, http.StatusBadRequest)
	}
	return common.SystemResponse(ec, res, nil, http.StatusOK)
}

func (sc *ShopController) UpsertCart(ec echo.Context) error {
	var (
		request shop.UpsertCartRequest
	)

	if err := ec.Bind(&request); err != nil {
		return common.SystemResponse(ec, nil, err, http.StatusBadRequest)
	}

	if err := ec.Validate(request); err != nil {
		return common.SystemResponse(ec, "invalid request", err, http.StatusBadRequest)
	}

	err := sc.shopInterface.UpsertCart(request.UserID, shop.Cart{
		SKU: request.SKU,
		Qty: request.Qty,
	})
	if err != nil {
		return common.SystemResponse(ec, nil, err, http.StatusBadRequest)
	}
	success := common.GeneralResponse{Success: true}
	return common.SystemResponse(ec, success, nil, http.StatusOK)
}
