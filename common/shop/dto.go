package shop

type Item struct {
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Qty   int     `json:"qty"`
}

type Cart struct {
	SKU    string     `json:"sku"`
	Qty    int        `json:"sku"`
	Status CartStatus `json:"status"`
}

type CheckoutRequest struct {
	Items  map[string]int `json:"items" validate:"required"`
	UserID int64          `json:"user_id" validate:"required"`
}

type CheckoutResponse struct {
	Items      []Item  `json:"items"`
	TotalPrice float32 `json:"total_price"`
}

type UpsertCartRequest struct {
	SKU    string `json:"sku" validate:"required"`
	Qty    int    `json:"qty" validate:"required"`
	UserID int64  `json:"user_id" validate:"required"`
}

type CartStatus string

const (
	CartStatusInCart CartStatus = "IN_CART"
	CartStatusDone   CartStatus = "DONE"
)
