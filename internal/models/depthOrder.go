package models

type DepthOrder struct {
	Price   float64 `json:"price"`
	BaseQty float64 `json:"base_qty"`
}

type OrderBook struct {
	ID       int64        `json:"id"`
	Exchange string       `json:"exchange"`
	Pair     string       `json:"pair"`
	Asks     []DepthOrder `json:"asks"`
	Bids     []DepthOrder `json:"bids"`
}
