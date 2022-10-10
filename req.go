package zoodmall_sdk

type ReqProductsFind struct {
	ProductID        ProductID `json:"productId"`        // Product ID.
	ProductIDs       string    `json:"productIds"`       // Comma(,) separated Product IDs
	SPU              string    `json:"spu"`              // Product SPU
	SKUs             string    `json:"skus"`             // Comma(,) separated Product SPUs
	Status           string    `json:"status"`           // Sale status (1: Yes 2: No)
	CategoryID       int64     `json:"categoryId"`       // Category of product
	NameSearch       string    `json:"nameSearch"`       // Keywords for product names
	DefaultPriceFrom float64   `json:"defaultPriceFrom"` // Minimum Price(>=)
	DefaultPriceTo   float64   `json:"defaultPriceTo"`   // Maximum Price(<=)
	Page             uint      `json:"page"`             // Current Page (Minimum: 1)
	RowsPerPage      uint      `json:"rowsPerPage"`      // Per Page products (Minimum: 1, Maximum: 20)
}

type ReqProductCreate struct {
	SPU          string                     `json:"spu"` // Must be unique
	Name         string                     `json:"name"`
	DefaultPrice float64                    `json:"defaultPrice"`
	Picture      []string                   `json:"picture"` // At least 1 image required (Maximum: 10, Minimum: 1)
	CategoryID   int64                      `json:"categoryId"`
	PropertyName string                     `json:"propertyName"` // Required if product is multi variant (i.e. property length > 1)
	Properties   []ReqProductCreateProperty `json:"property"`
	SProperties  []SProperty                `json:"sproperty"`
	Brand        string                     `json:"brand"`
	ShippingFee  float64                    `json:"shippingFee"`
	Description  string                     `json:"description"`
	Height       float32                    `json:"productHeight"`
	Lenght       float32                    `json:"productLenght"`
	Width        float32                    `json:"productWidth"`
	Weight       float32                    `json:"productWeight"`
	Tags         string                     `json:"tags"` // Comma(,) separated product tags
	CrossedPrice float64                    `json:"crossedPrice"`
	SizeGuideUrl string                     `json:"sizeGuideUrl"`
	OnSale       uint                       `json:"onSale"` // Sale status (1: Yes 2: No)
}

type ReqProductUpdate struct {
	ID              ProductID                  `json:"productId"`
	SPU             string                     `json:"spu"` // Required if productId is empty.
	Name            string                     `json:"name"`
	NameLang        map[string]string          `json:"nameLang"`
	Brand           string                     `json:"brand"`
	DefaultPrice    float64                    `json:"defaultPrice"`
	Status          string                     `json:"status"`
	Picture         []string                   `json:"picture"` // Maximum: 10, Minimum: 1
	CategoryID      int64                      `json:"categoryId"`
	PropertyName    string                     `json:"propertyName"` // Required if product is multi variant (i.e. property length > 1)
	Properties      []ReqProductUpdateProperty `json:"property"`
	SProperties     []SProperty                `json:"sproperty"`
	SPropertyLang   SPropertyLang              `json:"spropertyLang"`
	SKUType         uint                       `json:"skuType"`
	Description     string                     `json:"description"`
	DescriptionLang map[string]string          `json:"descriptionLang"`
	Tags            string                     `json:"tags"`     // Comma(,) separated product tags
	TagsLang        map[string]string          `json:"tagsLang"` // Comma(,) separated product other language tags
	CrossedPrice    float64                    `json:"crossedPrice"`
	SizeGuideUrl    string                     `json:"sizeGuideUrl"`
	ShippingFee     float64                    `json:"shippingFee"`
	OnSale          uint                       `json:"onSale"` // Sale status (1: Yes 2: No)
}

type ReqVariantCreate struct {
	ProductID     ProductID `json:"productId"`
	SPU           string    `json:"spu"` // Required if productId is empty.
	SKU           string    `json:"sku"`
	PropertyValue string    `json:"propertyValue"`
	Price         float64   `json:"price"`
	Quantity      uint      `json:"quantity"`
	PropertyImage string    `json:"propertyImage"`
	OnSale        uint      `json:"onSale"` // Sale status (1: Yes 2: No)
}

type ReqVariantUpdate struct {
	ProductID     ProductID `json:"productId"`
	SPU           string    `json:"spu"` // Required if productId is empty.
	SKU           string    `json:"sku"`
	PropertyValue string    `json:"propertyValue"`
	Price         float64   `json:"price"`
	Quantity      uint      `json:"quantity"`
	PropertyImage string    `json:"propertyImage"`
	OnSale        uint      `json:"onSale"` // Sale status (1: Yes 2: No)
}

type ReqProductCreateProperty struct {
	SKU      string  `json:"sku"` // Must be unique in same product.
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
	Value    string  `json:"propertyValue"`
}

type ReqProductUpdateProperty struct {
	SKU      string  `json:"sku"` // Must be unique in same product.
	Value    string  `json:"propertyValue"`
	Status   string  `json:"status"`
	Image    string  `json:"propertyImage"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
	OnSale   uint    `json:"onSale"` // Sale status (1: Yes 2: No)
}
