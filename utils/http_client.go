package utils

import (
	"crypto/tls"
	"net/http"
	"time"
)

func NewHttpClient(timeOut time.Duration, ignoreSSL bool) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: ignoreSSL,
		},
	}
	return &http.Client{
		Timeout:   timeOut,
		Transport: transport,
	}
}
