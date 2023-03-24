package hh

import (
	"net/http"
	"net/url"
)

type AppInfo struct {
	AuthType      string `json:"auth_type"`
	IsAdmin       bool   `json:"is_admin"`
	IsApplicant   bool   `json:"is_applicant"`
	IsApplication bool   `json:"is_application"`
	IsEmployer    bool   `json:"is_employer"`
}

func (c *Client) Me() (*AppInfo, error) {
	rel := &url.URL{Path: "/me"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res := AppInfo{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
