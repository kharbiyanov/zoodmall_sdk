package structures

type Product struct {
	ID              int64             `json:"productId"`
	SPU             string            `json:"spu"`
	Name            string            `json:"name"`
	NameLang        map[string]string `json:"nameLang"`
	DefaultPrice    float64           `json:"defaultPrice"`
	Picture         []string          `json:"picture"`
	CategoryID      int64             `json:"categoryId"`
	PropertyName    string            `json:"propertyName"`
	Properties      []Property        `json:"property"`
	SProperties     []SProperty       `json:"sproperty"`
	SPropertyLang   []SPropertyLang   `json:"spropertyLang"`
	SKUType         uint              `json:"skuType"`
	Currency        string            `json:"currency"`
	MarketCodes     string            `json:"marketCodes"`
	Description     string            `json:"description"`
	DescriptionLang map[string]string `json:"descriptionLang"`
	DeclaredValue   float64           `json:"declaredValue"`
	HSCode          string            `json:"hsCode"`
	HasBattery      uint              `json:"hasBattery"`
	HasPowder       uint              `json:"hasPowder"`
	HasLiquid       uint              `json:"hasLiquid"`
	Height          float32           `json:"productHeight"`
	Lenght          float32           `json:"productLenght"`
	Width           float32           `json:"productWidth"`
	Weight          float32           `json:"productWeight"`
	Tags            []string          `json:"tags"`
	TagsLang        map[string]string `json:"tagsLang"`
	CrossedPrice    float64           `json:"crossedPrice"`
	SizeGuideUrl    string            `json:"sizeGuideUrl"`
	ShippingFeeKZ   string            `json:"shippingFee_KZ"`
	ShippingFeeRU   string            `json:"shippingFee_RU"`
}

type ProductList []Product
