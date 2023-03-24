package hh

import (
	"net/http"
	"net/url"
)

// Get information about application
// API Documentation - https://api.hh.ru/openapi/redoc#tag/Informaciya-o-prilozhenii
func (c *Client) Me() (*AppInfo, error) {
	relURL := &url.URL{
		Path: "/me",
	}

	fullURL := c.BaseURL.ResolveReference(relURL)

	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		return nil, err
	}

	res := AppInfo{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type AppInfo struct {
	AuthType      string `json:"auth_type"`
	IsAdmin       bool   `json:"is_admin"`
	IsApplicant   bool   `json:"is_applicant"`
	IsApplication bool   `json:"is_application"`
	IsEmployer    bool   `json:"is_employer"`
}
