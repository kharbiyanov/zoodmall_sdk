package structures

type ResCreateProduct struct {
	ID  string `json:"productId"` // Product ID of specified SKU
	SKU string `json:"sku"`
}

type ResFindProducts struct {
	Products ProductList `json:"products"`
}
