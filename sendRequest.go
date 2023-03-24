package hh

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type errorItem struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type errorResponse struct {
	Description string      `json:"description"`
	OauthError  string      `json:"oauth_error"`
	Errors      []errorItem `json:"errors"`
	RequestId   string      `json:"request_id"`
}

/*
*

	Since all API endpoints act in the same manner, helper function sendRequest is created to avoid code duplication.
	It will set common headers (content type, auth header), make request, check for errors, parse response.
*/
func (c *Client) sendRequest(req *http.Request, fullResponse interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AppAccessToken))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Error handling
	// We're considering status codes < 200 and >= 400 as errors and parse response into errorResponse.
	// It depends on the API design though, your API may handle errors differently.
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Description)
		}

		return fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}

	// Trying to write a successful answer into pointer fullResponse
	if err = json.NewDecoder(resp.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}
