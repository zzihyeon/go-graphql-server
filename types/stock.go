package types

type StockType struct {
	Name   string   `json:"name"`
	Volume string   `json:"volume"`
	Reason []string `json:"reason"`
}
