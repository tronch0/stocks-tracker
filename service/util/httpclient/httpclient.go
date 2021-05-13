package httpclient

import (
	"errors"
	"fmt"
	"github.com/ugorji/go/codec"
	"io/ioutil"
	"log"
	"net/http"
)

var authErr = errors.New("authentication failed. request is not authorized")
var jsonHandle codec.JsonHandle

type HttpClient struct {
	authToken string
}

func New() *HttpClient {
	return &HttpClient{}
}

func (c *HttpClient) SetAuthToken(token string) {
	c.authToken = token
}

func (c *HttpClient) SendGetRequest(url string, queryParam map[string]string) (body []byte, err error) {
	req, err := c.createGetRequest(url, queryParam)
	if err != nil {
		return nil, err
	}
	log.Println("sending GET request to url " + url)
	res, err := c.sendRequest(req)
	if err != nil {
		return nil,  err
	}
	defer res.Body.Close()

	if err := checkForSuccessfulResponse(res.StatusCode); err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (c *HttpClient) createGetRequest(url string, queryParam map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if len(c.authToken) > 0 {
		req.Header.Set("authorization", "Bearer "+c.authToken)
	}

	if len(queryParam) > 0 {
		q := req.URL.Query()

		for key, value := range queryParam {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	return req, nil
}
func (c *HttpClient) sendRequest(req *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(req)
}

func checkForSuccessfulResponse(httpStatus int) (err error) {
	if httpStatus < 200 || httpStatus > 299 {
		if httpStatus == 401 {
			err = authErr
		} else {
			err = fmt.Errorf("http response with status code %d", httpStatus)
		}

		log.Printf("error on http-response, err: %v", err)
		return err
	}

	return nil
}