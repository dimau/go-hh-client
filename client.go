package hh

import (
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL        *url.URL // https://api.hh.ru
	UserAgent      string   // MyApp/1.0 (my-app-feedback@example.com)
	AppAccessToken string   // Access token for application registered in hh.ru
	HTTPClient     *http.Client
}

func NewClient(baseUrl *url.URL, userAgent string, appAccessToken string) *Client {
	return &Client{
		BaseURL:        baseUrl,
		UserAgent:      userAgent,
		AppAccessToken: appAccessToken,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
