package structures

type ReqFindProducts struct {
	ProductID        int64   `json:"productId"`        // Product ID.
	ProductIDs       string  `json:"productIds"`       // Comma(,) separated Product IDs
	SPU              string  `json:"spu"`              // Product SPU
	SKUs             string  `json:"skus"`             // Comma(,) separated Product SPUs
	Status           string  `json:"status"`           // Sale status (1: Yes 2: No)
	CategoryID       int64   `json:"categoryId"`       // Category of product
	NameSearch       string  `json:"nameSearch"`       // Keywords for product names
	DefaultPriceFrom float64 `json:"defaultPriceFrom"` // Minimum Price(>=)
	DefaultPriceTo   float64 `json:"defaultPriceTo"`   // Maximum Price(<=)
	Page             uint    `json:"page"`             // Current Page (Minimum: 1)
	RowsPerPage      uint    `json:"rowsPerPage"`      // Per Page products (Minimum: 1, Maximum: 20)
}

type ReqCreateProduct struct {
	SPU          string                      `json:"spu"` // Must be unique
	Name         string                      `json:"name"`
	DefaultPrice float64                     `json:"defaultPrice"`
	Picture      []string                    `json:"picture"` // At least 1 image required (Maximum: 10, Minimum: 1)
	CategoryID   int64                       `json:"categoryId"`
	PropertyName string                      `json:"propertyName"` // Required if product is multi variant (i.e. property length > 1)
	Properties   []ReqCreateProductProperty  `json:"property"`
	SProperties  []ReqCreateProductSProperty `json:"sproperty"`
	Brand        string                      `json:"brand"`
	ShippingFee  float64                     `json:"shippingFee"`
	Description  string                      `json:"description"`
	Height       float32                     `json:"productHeight"`
	Lenght       float32                     `json:"productLenght"`
	Width        float32                     `json:"productWidth"`
	Weight       float32                     `json:"productWeight"`
	Tags         string                      `json:"tags"` // Comma(,) separated product tags
	CrossedPrice float64                     `json:"crossedPrice"`
	SizeGuideUrl string                      `json:"sizeGuideUrl"`
	OnSale       uint                        `json:"onSale"` // Sale status (1: Yes 2: No)
}

type ReqCreateProductProperty struct {
	SKU      string  `json:"sku"` // Must be unique in same product.
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
	Value    string  `json:"propertyValue"`
}

type ReqCreateProductSProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
