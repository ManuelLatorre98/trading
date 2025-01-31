package model

import "time"

type DollarModel struct {
	DolarId   int32     `json:"id"`
	DolarDate time.Time `json:"dolarDate"`
	DolarType string    `json:"dolarType"`
	BuyPrice  float32   `json:"buyPrice"`
	SellPrice float32   `json:"sellPrice"`
}

type DollarInput struct {
	DolarType string  `json:"dolarType"`
	BuyPrice  float32 `json:"buyPrice"`
	SellPrice float32 `json:"sellPrice"`
}
