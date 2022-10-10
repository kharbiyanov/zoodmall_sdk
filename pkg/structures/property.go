package structures

type Property struct {
	SKU      string  `json:"sku"`
	Value    string  `json:"propertyValue"`
	Image    string  `json:"propertyImage"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
	OnSale   uint    `json:"onSale"`
}

type SProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SPropertyLang map[string][]SProperty
