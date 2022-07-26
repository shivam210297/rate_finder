package models

import "time"

type Rates struct {
	Time    time.Time `json:"time"`
	AssetID string    `json:"asset_id_quote"`
	Price   float64   `json:"rate"`
}
type CoinInfo struct {
	AssetIDBase string  `json:"asset_id_base"`
	RateInfo    []Rates `json:"rates"`
}
