package utils

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	netHttp *http.Client
}

func NewHttpClient(timeOut time.Duration, ignoreSSL bool) *HttpClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: ignoreSSL,
		},
	}
	return &HttpClient{
		netHttp: &http.Client{
			Timeout:   timeOut,
			Transport: transport,
		},
	}
}

func (h *HttpClient) Request(method, url string, body io.Reader) (statusCode int, contentLength int64, responseBody string) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	request.Header.Add("User-Agent", GetRandomUa())
	response, err := h.netHttp.Do(request)
	if err != nil {
		return
	}
	b, _ := ioutil.ReadAll(response.Body)
	return response.StatusCode, response.ContentLength, string(b)

}
