package hh

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestGetVacanciesWithPeriod(t *testing.T) {
	appAccessToken := os.Getenv("HH_API_KEY")

	client := NewClient(
		&url.URL{
			Scheme: "https",
			Host:   "api.hh.ru",
		},
		"dimau-app/1.0 (dimau777@gmail.com)",
		appAccessToken)

	// Options for getting React vacancies
	options := &OptionsForGetVacancies{
		Text:         "react",
		SearchField:  "name",
		Period:       2,
		ItemsPerPage: 6,
		PageNumber:   0,
		DateFrom:     nil,
		OrderBy:      "",
	}

	// Get vacancies
	vacancies, err := client.GetVacancies(options)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, vacancies, "expecting non-nil result")
}

func TestGetVacanciesWithFromPublishTime(t *testing.T) {
	appAccessToken := os.Getenv("HH_API_KEY")

	client := NewClient(
		&url.URL{
			Scheme: "https",
			Host:   "api.hh.ru",
		},
		"dimau-app/1.0 (dimau777@gmail.com)",
		appAccessToken)

	// Options for getting React vacancies
	fromPublishTime := time.Now().Truncate(10 * time.Hour)
	options := &OptionsForGetVacancies{
		Text:         "react",
		SearchField:  "name",
		Period:       0,
		ItemsPerPage: 6,
		PageNumber:   0,
		DateFrom:     &fromPublishTime,
		OrderBy:      "publication_time",
	}

	// Get vacancies
	vacancies, err := client.GetVacancies(options)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, vacancies, "expecting non-nil result")
}
