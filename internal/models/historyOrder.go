package models

import "time"

type HistoryOrder struct {
	ClientName          string    `json:"client_name"`
	ExchangeName        string    `json:"exchange_name"`
	Label               string    `json:"label"`
	Pair                string    `json:"pair"`
	Side                string    `json:"side"`
	Type                string    `json:"type"`
	BaseQty             float64   `json:"base_qty"`
	Price               float64   `json:"price"`
	AlgorithmNamePlaced string    `json:"algorithm_name_placed"`
	LowestSellPrice     float64   `json:"lowest_sell_price"`
	HighestBuyPrice     float64   `json:"highest_buy_price"`
	CommissionQuoteQty  float64   `json:"commission_quote_qty"`
	TimePlaced          time.Time `json:"time_placed"`
}
