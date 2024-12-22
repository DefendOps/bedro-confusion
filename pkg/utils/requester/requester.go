package requester

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var (
	ErrPageNotFound     = errors.New("404 Not Found")
)

type HTTPRequest struct{
	BaseURL    	string
	Endpoint    string
	Method		string
	IsJson		bool
	Body		interface{}
	Headers		map[string]string
}

type HTTPRequester struct {}

func (hr *HTTPRequester) PerformRequest(request HTTPRequest) (*http.Response, error){
	var body []byte
	
	if request.IsJson {
		if requestBodyBytes, ok := request.Body.([]byte); ok {
			body = requestBodyBytes
		} else {
			body, _ = json.Marshal(request.Body)
		}
	} else {
		if requestBodyBytes, ok := request.Body.([]byte); ok {
			body = requestBodyBytes
		} else {
			body = []byte(request.Body.(string))
		}
	}

	req, err := http.NewRequest(request.Method, request.BaseURL + request.Endpoint, bytes.NewBuffer(body))

	if err != nil {
		return &http.Response{}, err
	}
	
	req.Header.Set("User-Agent", "BEDRO-CONFUSER (A tool i'm writing to have some fun -> https://bedro.defendops.com/)")

	if request.IsJson{
		req.Header.Set("Content-Type", "application/json")
	}
	
	if len(request.Headers) > 0{
		for key, value := range request.Headers{
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	if res.StatusCode != http.StatusOK {
		return &http.Response{}, ErrPageNotFound
	}

	return res, nil
}