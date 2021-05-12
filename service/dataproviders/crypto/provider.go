package crypto

import (
	"fmt"
	"github.com/ugorji/go/codec"
	"stocks_tracker/service/util/httpclient"
	"time"
)

var jsonHandle codec.JsonHandle

const (
	SERVICE_ADDRESS = "https://api.coingecko.com/api/v3"
)


type CryptoProvider struct {
	serviceClient *httpclient.HttpClient
}

func New() *CryptoProvider {
	return &CryptoProvider{
		serviceClient: httpclient.New(),
	}
}

func (cp *CryptoProvider) GetQuote(symbol string) (float64, error) {
	param := make(map[string]string)
	param["ids"] = symbol
	param["vs_currencies"] = "usd"
	res := make(map[string]map[string]float64)

	err := cp.getQuoteFromExternalProvider(SERVICE_ADDRESS + "/simple/price", param, &res)
	if err != nil {
		return 0, err
	}

	if coin, isCoinExist := res[param["ids"]]; isCoinExist {
		if price, isPriceExist := coin[param["vs_currencies"]];isPriceExist {
			return price, nil
		}
	}

	return 0, fmt.Errorf("coin \"%s\" or price in \"%s\" is not avilable",param["ids"],param["vs_currencies"] )
}
func (cp *CryptoProvider) GetQuoteByDate(symbol string, date time.Time) (float64, error) {
	param := make(map[string]string)
	param["date"] = date.Format("02-01-2006")
	param["localization"] = "false"

	res := &CoinsIDHistory{}

	err := cp.getQuoteFromExternalProvider(SERVICE_ADDRESS + fmt.Sprintf("/coins/%s/history",symbol), param, res)
	if err != nil {
		return 0, err
	}

	return res.MarketData.CurrentPrice["usd"], nil
}


func (cp *CryptoProvider) getQuoteFromExternalProvider(address string, param map[string]string, res interface{}) error {
	resBytes, err := cp.serviceClient.SendGetRequest(address, param)
	if err != nil {
		return err
	}

	err = UnmarshalResponse(resBytes,&res)
	if err != nil {
		return err
	}

	return nil
}

func UnmarshalResponse(body []byte, targetRequestObject interface{}) error {
	return codec.NewDecoderBytes(body, &jsonHandle).Decode(targetRequestObject)
}