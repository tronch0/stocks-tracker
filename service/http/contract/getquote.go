package contract

type GetQuoteResponse struct {
	Price float64 `codec:"price"`
	Id    string  `codec:"id"`
}
