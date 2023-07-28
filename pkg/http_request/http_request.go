package http_request

import (
	"bytes"
	"net/http"
)

func DoGet(url string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := prepareRequest(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DoPost(url string, headers map[string]string, body []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := prepareRequest(http.MethodPost, url, headers, body)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func prepareRequest(method string, url string, headers map[string]string, body []byte) (*http.Request, error) {
	var req *http.Request
	var err error
	if method == http.MethodGet {
		req, err = http.NewRequest(http.MethodGet, url, nil)
	} else if method == http.MethodPost {
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	}
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req, nil
}
