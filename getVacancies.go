package hh

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type OptionsForGetVacancies struct {
	Text         string     // "react"
	SearchField  string     // "name"
	Period       int        // 1
	ItemsPerPage int        // 20
	PageNumber   int        // 0
	DateFrom     *time.Time //
	OrderBy      string     // "publication_time" / "salary_desc" ...
}

// Get all vacancies from hh.ru that are satisfy the specified options
// API Documentation - https://github.com/hhru/api/blob/master/docs/vacancies.md#поиск-по-вакансиям
func (c *Client) GetVacancies(options *OptionsForGetVacancies) (*Vacancies, error) {
	relURL := &url.URL{
		Path:     "/vacancies",
		RawQuery: fmt.Sprintf("text=%v&search_field=%v&period=%v&per_page=%v&page=%v&date_from=%v&order_by=%v", options.Text, options.SearchField, options.Period, options.ItemsPerPage, options.PageNumber, convertGoTimeToISO8601(options.DateFrom), options.OrderBy),
	}

	fullURL := c.BaseURL.ResolveReference(relURL)

	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		return nil, err
	}

	res := Vacancies{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type Vacancies struct {
	PerPage   int         `json:"per_page"`
	Page      int         `json:"page"`
	Pages     int         `json:"pages"`
	Found     int         `json:"found"`
	Clusters  interface{} `json:"clusters"`
	Arguments interface{} `json:"arguments"`
	Items     []Vacancy
}

type Vacancy struct {
	Salary                 Salary           `json:"salary"`
	Name                   string           `json:"name"`
	InsiderInterview       InsiderInterview `json:"insider_interview"`
	Area                   Area             `json:"area"`
	Url                    string           `json:"url"`
	PublishedAt            string           `json:"published_at"`
	Relations              []interface{}    `json:"relations"`
	Employer               Employer         `json:"employer"`
	Contacts               Contacts         `json:"contacts"`
	ResponseLetterRequired bool             `json:"response_letter_required"`
	Address                Address          `json:"address"`
	SortPointDistance      float64          `json:"sort_point_distance"`
	AlternateUrl           string           `json:"alternate_url"`
	ApplyAlternateUrl      string           `json:"apply_alternate_url"`
	Department             Department       `json:"department"`
	Type                   Type             `json:"type"`
	Id                     string           `json:"id"`
	HasTest                bool             `json:"has_test"`
	ResponseUrl            interface{}      `json:"response_url"`
	Snippet                Snippet          `json:"snippet"`
	Schedule               Schedule         `json:"schedule"`
	Counters               Counters         `json:"counters"`
}

type Salary struct {
	To       interface{} `json:"to"`
	From     int         `json:"from"`
	Currency string      `json:"currency"`
	Gross    bool        `json:"gross"`
}

type InsiderInterview struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type Area struct {
	Url  string `json:"url"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Employer struct {
	LogoUrls struct {
		Field1   string `json:"90"`
		Field2   string `json:"240"`
		Original string `json:"original"`
	} `json:"logo_urls"`
	Name         string `json:"name"`
	Url          string `json:"url"`
	AlternateUrl string `json:"alternate_url"`
	Id           string `json:"id"`
	Trusted      bool   `json:"trusted"`
}

type Phone struct {
	Country string      `json:"country"`
	City    string      `json:"city"`
	Number  string      `json:"number"`
	Comment interface{} `json:"comment"`
}

type Contacts struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Phones []Phone `json:"phones"`
}

type MetroStation struct {
	StationId   string  `json:"station_id"`
	StationName string  `json:"station_name"`
	LineId      string  `json:"line_id"`
	LineName    string  `json:"line_name"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}

type Address struct {
	City          string         `json:"city"`
	Street        string         `json:"street"`
	Building      string         `json:"building"`
	Description   string         `json:"description"`
	Lat           float64        `json:"lat"`
	Lng           float64        `json:"lng"`
	MetroStations []MetroStation `json:"metro_stations"`
}

type Department struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Type struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Snippet struct {
	Requirement    string `json:"requirement"`
	Responsibility string `json:"responsibility"`
}

type Schedule struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Counters struct {
	Responses int `json:"responses"`
}
