package api

type Product struct {
	Code     string    `json:"code"`
	Title    string    `json:"title"`
	Vendor   string    `json:"vendor"`
	BodyHtml string    `json:"body_html"`
	Variants []Variant `json:"variants"`
	Images   []Image   `json:"images"`
}

type Variant struct {
	Id                string `json:"id"`
	Title             string `json:"title"`
	Sku               string `json:"sku"`
	Available         bool   `json:"available"`
	InventoryQuantity int64  `json:"inventory_quantity"`
	Weight            Weight `json:"weight"`
}

type Weight struct {
	Value int64  `json:"value"`
	Unit  string `json:"weight_unit"`
}

type Image struct {
	Source    string `json:"src"`
	VariantID string `json:"variantId"`
}

type Error struct {
	Message string `json:"message"`
}

type RequestProduct struct {
	Id       int              `json:"id"`
	Code     string           `json:"code"`
	Title    string           `json:"title"`
	Vendor   string           `json:"vendor"`
	BodyHtml string           `json:"body_html"`
	Variants []RequestVariant `json:"variants"`
}

type RequestVariant struct {
	Id                int     `json:"id"`
	Title             string  `json:"title"`
	Sku               string  `json:"sku"`
	Position          int     `json:"position"`
	Available         bool    `json:"available"`
	InventoryQuantity int64   `json:"inventory_quantity"`
	Weight            int64   `json:"weight"`
	WeightUnit        string  `json:"weight_unit"`
	Images            []Image `json:"images"`
}

type Inventory struct {
	ProductId string `json:"productId"`
	VariantId string `json:"variantId"`
	Stock     int64  `json:"stock"`
}
