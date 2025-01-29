package model

import "time"

type DolarModel struct {
	DolarDate     time.Time `json:"dolarDate"`
	PriceBlue     float32   `json:"priceBlue"`
	PriceOfficial float32   `json:"priceOfficial"`
	PriceMep      float32   `json:"priceMep"`
}

type DolarInput struct {
	PriceBlue     float32 `json:"priceBlue"`
	PriceOfficial float32 `json:"priceOfficial"`
	PriceMep      float32 `json:"priceMep"`
}
