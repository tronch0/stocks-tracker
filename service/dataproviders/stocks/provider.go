package stocks

import (
	"github.com/ugorji/go/codec"
	"stocks_tracker/service/util/httpclient"
)

var jsonHandle codec.JsonHandle

const (
	SERVICE_ADDRESS = "http://api.marketstack.com/v1/eod"
	AUTH_TOKEN = "5c2aa5bec3c6d7187df18ed14e515b4b"
)

type StocksProvider struct {
	serviceClient *httpclient.HttpClient
}

func New() *StocksProvider {
	return &StocksProvider{
		serviceClient: httpclient.New(AUTH_TOKEN),
	}
}

func (sp *StocksProvider) GetQuote(symbol string) (float64, error) {
	param := make(map[string]string)
	param["symbols"] = symbol
	param["access_key"] = AUTH_TOKEN
	param["limit"] = "10"

	return sp.getQuoteFromExternalProvider(SERVICE_ADDRESS + "/latest", param)
}
func (sp *StocksProvider) GetQuoteByDate(symbol, date string) (float64, error) {
	param := make(map[string]string)
	param["symbols"] = symbol
	param["date_from"] =  date
	param["date_to"] =  date
	param["access_key"] = AUTH_TOKEN
	param["limit"] = "10"

	return sp.getQuoteFromExternalProvider(SERVICE_ADDRESS, param)
}

func (sp *StocksProvider) getQuoteFromExternalProvider(address string, param map[string]string) (float64, error) {
	resBytes, err := sp.serviceClient.SendGetRequest(address, param)
	if err != nil {
		return 0, err
	}
	res := &HttpResponseQuote{}
	err = UnmarshalResponse(resBytes,res)
	if err != nil {
		return 0, err
	}

	if len(res.Data) == 0 {
		return 0, nil
	}

	return res.Data[0].Close, nil
}

func UnmarshalResponse(body []byte, targetRequestObject interface{}) error {
	return codec.NewDecoderBytes(body, &jsonHandle).Decode(targetRequestObject)
}