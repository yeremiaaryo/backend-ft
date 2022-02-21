// Code generated by MockGen. DO NOT EDIT.
// Source: application/shop/application.go

// Package mock_shop_application is a generated GoMock package.
package mock_shop_application

import (
	shop "backend-ft/common/shop"
	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	reflect "reflect"
)

// MockShopApplication is a mock of ShopApplication interface
type MockShopApplication struct {
	ctrl     *gomock.Controller
	recorder *MockShopApplicationMockRecorder
}

// MockShopApplicationMockRecorder is the mock recorder for MockShopApplication
type MockShopApplicationMockRecorder struct {
	mock *MockShopApplication
}

// NewMockShopApplication creates a new mock instance
func NewMockShopApplication(ctrl *gomock.Controller) *MockShopApplication {
	mock := &MockShopApplication{ctrl: ctrl}
	mock.recorder = &MockShopApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockShopApplication) EXPECT() *MockShopApplicationMockRecorder {
	return m.recorder
}

// GetItemByMultipleSKU mocks base method
func (m *MockShopApplication) GetItemByMultipleSKU(sku []string, tx *sqlx.Tx) ([]shop.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemByMultipleSKU", sku, tx)
	ret0, _ := ret[0].([]shop.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemByMultipleSKU indicates an expected call of GetItemByMultipleSKU
func (mr *MockShopApplicationMockRecorder) GetItemByMultipleSKU(sku, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemByMultipleSKU", reflect.TypeOf((*MockShopApplication)(nil).GetItemByMultipleSKU), sku, tx)
}

// GetItemBySKU mocks base method
func (m *MockShopApplication) GetItemBySKU(sku string) (*shop.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemBySKU", sku)
	ret0, _ := ret[0].(*shop.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemBySKU indicates an expected call of GetItemBySKU
func (mr *MockShopApplicationMockRecorder) GetItemBySKU(sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemBySKU", reflect.TypeOf((*MockShopApplication)(nil).GetItemBySKU), sku)
}

// InsertCheckout mocks base method
func (m *MockShopApplication) InsertCheckout(userId int64, items []shop.Item, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCheckout", userId, items, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCheckout indicates an expected call of InsertCheckout
func (mr *MockShopApplicationMockRecorder) InsertCheckout(userId, items, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCheckout", reflect.TypeOf((*MockShopApplication)(nil).InsertCheckout), userId, items, tx)
}

// UpdateItemQty mocks base method
func (m *MockShopApplication) UpdateItemQty(items []shop.Item, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemQty", items, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItemQty indicates an expected call of UpdateItemQty
func (mr *MockShopApplicationMockRecorder) UpdateItemQty(items, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemQty", reflect.TypeOf((*MockShopApplication)(nil).UpdateItemQty), items, tx)
}

// GetCartByUserIDAndStatusInCart mocks base method
func (m *MockShopApplication) GetCartByUserIDAndStatusInCart(userID int64) ([]shop.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartByUserIDAndStatusInCart", userID)
	ret0, _ := ret[0].([]shop.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartByUserIDAndStatusInCart indicates an expected call of GetCartByUserIDAndStatusInCart
func (mr *MockShopApplicationMockRecorder) GetCartByUserIDAndStatusInCart(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartByUserIDAndStatusInCart", reflect.TypeOf((*MockShopApplication)(nil).GetCartByUserIDAndStatusInCart), userID)
}

// GetCartByUserIDAndSKU mocks base method
func (m *MockShopApplication) GetCartByUserIDAndSKU(userID int64, cart shop.Cart, tx *sqlx.Tx) (*shop.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartByUserIDAndSKU", userID, cart, tx)
	ret0, _ := ret[0].(*shop.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartByUserIDAndSKU indicates an expected call of GetCartByUserIDAndSKU
func (mr *MockShopApplicationMockRecorder) GetCartByUserIDAndSKU(userID, cart, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartByUserIDAndSKU", reflect.TypeOf((*MockShopApplication)(nil).GetCartByUserIDAndSKU), userID, cart, tx)
}

// InsertCart mocks base method
func (m *MockShopApplication) InsertCart(userId int64, cart shop.Cart, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCart", userId, cart, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCart indicates an expected call of InsertCart
func (mr *MockShopApplicationMockRecorder) InsertCart(userId, cart, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCart", reflect.TypeOf((*MockShopApplication)(nil).InsertCart), userId, cart, tx)
}

// UpdateCart mocks base method
func (m *MockShopApplication) UpdateCart(userId int64, cart shop.Cart, status shop.CartStatus, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", userId, cart, status, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCart indicates an expected call of UpdateCart
func (mr *MockShopApplicationMockRecorder) UpdateCart(userId, cart, status, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockShopApplication)(nil).UpdateCart), userId, cart, status, tx)
}

// UpdateCartStatus mocks base method
func (m *MockShopApplication) UpdateCartStatus(userId int64, sku string, status shop.CartStatus, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartStatus", userId, sku, status, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCartStatus indicates an expected call of UpdateCartStatus
func (mr *MockShopApplicationMockRecorder) UpdateCartStatus(userId, sku, status, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartStatus", reflect.TypeOf((*MockShopApplication)(nil).UpdateCartStatus), userId, sku, status, tx)
}
