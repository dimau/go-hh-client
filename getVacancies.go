package hh

import (
	"net/http"
	"net/url"
)

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

type Contacts struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phones []struct {
		Country string      `json:"country"`
		City    string      `json:"city"`
		Number  string      `json:"number"`
		Comment interface{} `json:"comment"`
	} `json:"phones"`
}

type Address struct {
	City          string  `json:"city"`
	Street        string  `json:"street"`
	Building      string  `json:"building"`
	Description   string  `json:"description"`
	Lat           float64 `json:"lat"`
	Lng           float64 `json:"lng"`
	MetroStations []struct {
		StationId   string  `json:"station_id"`
		StationName string  `json:"station_name"`
		LineId      string  `json:"line_id"`
		LineName    string  `json:"line_name"`
		Lat         float64 `json:"lat"`
		Lng         float64 `json:"lng"`
	} `json:"metro_stations"`
}

type Department struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
	Type                   struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Id          string      `json:"id"`
	HasTest     bool        `json:"has_test"`
	ResponseUrl interface{} `json:"response_url"`
	Snippet     struct {
		Requirement    string `json:"requirement"`
		Responsibility string `json:"responsibility"`
	} `json:"snippet"`
	Schedule struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"schedule"`
	Counters struct {
		Responses int `json:"responses"`
	} `json:"counters"`
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

func (c *Client) GetVacancies() (*Vacancies, error) {
	rel := &url.URL{
		Path:     "/vacancies",
		RawQuery: "text=react&search_field=name&period=1",
	}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res := Vacancies{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

	//req.Header.Set("Authorization", "Bearer "+c.AppAccessToken)
	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("User-Agent", c.UserAgent)
	//
	//resp, err := c.HTTPClient.Do(req)
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
	//	var errRes errorResponse
	//	if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
	//		return nil, errors.New(errRes.Description)
	//	}
	//
	//	return nil, fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	//}
	//
	//var vacancies *Vacancies
	//var p = make([]byte, 1000)
	//_, err = resp.Body.Read(p)
	//fmt.Println(string(p))
	//err = json.NewDecoder(resp.Body).Decode(&vacancies)
	//return vacancies, err
}
