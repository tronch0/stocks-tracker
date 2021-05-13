package dataproviders

import "time"

type Provider interface {
	GetQuote(id string) (float64, error)
	GetQuoteByDate(id string, date time.Time) (float64, error)
}

