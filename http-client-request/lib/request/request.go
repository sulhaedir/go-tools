package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var client http.Client

type Request struct {
	URL     string
	Headers map[string]string
	Method  string
}

func init() {
	client = http.Client{
		Timeout: time.Second * 5,
	}
}

func NewRequest() *Request {
	return &Request{
		Headers: map[string]string{},
	}
}

func (r *Request) ValidateUrl() error {
	if r.URL == "" || len(r.URL) == 0 {
		fmt.Errorf("Url is not Valid")
	}
	return nil
}
func (r *Request) ValidateMethod() error {
	if r.Method != "GET" || r.Method != "POST" {
		fmt.Errorf("Method is not valid")
	}
	return nil
}

func (r *Request) Get(u *url.URL) (*http.Request, error) {
	resp, err := http.NewRequest(r.Method, r.URL, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Request) Exec() (*http.Response, []byte, error) {

	var request *http.Request
	var response *http.Response

	if err := r.ValidateUrl(); err != nil {
		return response, nil, err
	}

	if err := r.ValidateMethod(); err != nil {
		return response, nil, err
	}
	u, err := url.Parse(r.URL)
	if r.Method == "GET" {
		request, err = r.Get(u)
	}
	for key, value := range r.Headers {
		request.Header.Set(key, value)
	}
	request.Close = true

	response, err = client.Do(request)
	if err != nil {
		return response, nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response, nil, err
	}
	return response, body, nil
}
