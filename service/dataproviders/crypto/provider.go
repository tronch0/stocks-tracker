package crypto

import (
	"stocks_tracker/service/util/httpclient"
	"time"
)

const (
	SERVICE_ADDRESS = "127.0.0.1:"
	AUTH_TOKEN = ""
)


type CryptoProvider struct {
	serviceClient *httpclient.HttpClient
}

func New() *CryptoProvider {
	return &CryptoProvider{
		serviceClient: httpclient.New(AUTH_TOKEN),
	}
}

func (cp *CryptoProvider) GetQuote(symbol string) (float64, error) {

	return 0
}
func (cp *CryptoProvider) GetQuoteByDate(symbol string, date string) (float64, error) {

	return 0
}


func (cp *CryptoProvider) getQuoteFromExternalProvider(symbol string, date time.Time) (float64, error) {

}