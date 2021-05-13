package stocks

import (
	"github.com/ugorji/go/codec"
	"stocks_tracker/service/util/httpclient"
	"time"
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
	res := &StocksProvider{
		serviceClient: httpclient.New(),
	}

	res.serviceClient.SetAuthToken(AUTH_TOKEN)

	return res
}

func (sp *StocksProvider) GetQuote(id string) (float64, error) {
	param := make(map[string]string)
	param["symbols"] = id
	param["access_key"] = AUTH_TOKEN
	param["limit"] = "10"

	return sp.getQuoteFromExternalProvider(SERVICE_ADDRESS + "/latest", param)
}
func (sp *StocksProvider) GetQuoteByDate(id string, date time.Time) (float64, error) {

	param := make(map[string]string)
	param["symbols"] = id
	param["date_from"] =  date.Format( "2006-01-02")
	param["date_to"] =  param["date_from"]
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