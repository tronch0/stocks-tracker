package httpclient

import (
	"bytes"
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

func New(authToken string) *HttpClient {
	return &HttpClient{authToken: authToken}
}

func (c *HttpClient) SetAuthToken(token string) {
	log.Println("set auth-token to the http-client")
	c.authToken = token
}

func (c *HttpClient) IsAuthenticated() bool {
	return c.authToken != ""
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

func (c *HttpClient) SendPostRequest(reqBody interface{}, url string) (body []byte, httpStatus int, err error) {
	req, err := c.createPostRequest(reqBody, url)
	if err != nil {
		return nil, 0, err
	}
	log.Println("sending POST request to url " + url)
	res, err := c.sendRequest(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	return resBody, res.StatusCode, nil
}
func (c *HttpClient) createPostRequest(body interface{}, url string) (*http.Request, error) {
	jsonBody := []byte{}
	enc := codec.NewEncoderBytes(&jsonBody, &jsonHandle)
	err := enc.Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	if len(c.authToken) > 0 {
		req.Header.Set("authorization", "Bearer "+c.authToken)
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
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
