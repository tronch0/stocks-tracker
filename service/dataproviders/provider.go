package dataproviders

type Provider interface {
	GetQuote(symbol string) (float64, error)
	GetQuoteByDate(symbol string, date string) (float64, error)
}

