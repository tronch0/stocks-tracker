package stocks

import (
	"stocks_tracker/service/util/httpclient"
	"time"
)

const (
	SERVICE_ADDRESS = "https://alpha-vantage.p.rapidapi.com/query?function=GLOBAL_QUOTE&symbol=MSFT&datatype=json"
	AUTH_TOKEN = "2XM3XCP7GT9LNLOY"
)

type StocksProvider struct {
	serviceClient *httpclient.HttpClient
}

func New() *StocksProvider {
	return &StocksProvider{
		serviceClient: httpclient.New(AUTH_TOKEN),
	}
}

func (sp *StocksProvider) GetQuote(symbol string) float64 {

	return 0
}
func (sp *StocksProvider) GetQuoteByDate(symbol string, date time.Time) float64 {

	return 0
}


func (sp *StocksProvider) getQuoteFromExternalProvider(symbol string, date time.Time) float64 {

}