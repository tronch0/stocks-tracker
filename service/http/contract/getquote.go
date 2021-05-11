package contract

type GetQuoteResponse struct {
	Price float64 `json:"price"`
	Symbol string `json:"symbol"`
}
