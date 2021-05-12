package dataproviders

import "time"

type Provider interface {
	GetQuote(symbol string) (float64, error)
	GetQuoteByDate(symbol string, date time.Time) (float64, error)
}

