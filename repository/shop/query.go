package shop

var (
	getItemByMultipleSKU           = `SELECT sku, name, price, qty FROM items WHERE sku IN (?) FOR UPDATE`
	getItemBySKU                   = `SELECT sku, name, price, qty FROM items WHERE sku = ?`
	getCartByUserIDAndStatus       = `SELECT sku, qty FROM cart WHERE user_id = ? AND status = ?`
	getCartByUserIDAndSKUForUpdate = `SELECT sku, qty FROM cart WHERE user_id = ? AND sku = ? FOR UPDATE`
)

var (
	insertCart     = `INSERT INTO cart (user_id, sku, qty, status) VALUES (?, ?, ?, ?)`
	insertCheckout = `INSERT INTO checkout (user_id, sku, name, price, qty) VALUES (?, ?, ?, ?, ?)`
)

var (
	updateItemQty    = `UPDATE items SET qty = ? WHERE sku = ?`
	updateCart       = `UPDATE cart SET qty = ?, status = ? WHERE user_id = ? AND sku = ?`
	updateCartStatus = `UPDATE cart SET status = ? WHERE user_id = ? AND sku = ?`
)
