package libanemos

import (
	"testing"
	"time"

	"github.com/moznion/go-optional"
)

func TestWeatherWarningFilter(t *testing.T) {
	warnings := []WeatherWarning{
		{reported_at: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)},
		{reported_at: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC)},
		{reported_at: time.Date(2023, 1, 3, 10, 0, 0, 0, time.UTC)},
	}

	list := WeatherWarninglist{data: warnings}

	startTime, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2023-01-02T23:59:59Z")

	filterOptions := FilterOptions{
		StartTime: optional.Some(startTime),
		EndTime:   optional.Some(endTime),
	}

	filtered := list.Filter(filterOptions)

	if len(filtered.data) != 2 {
		t.Errorf("Expected 2 warnings, got %d", len(filtered.data))
	}
}

func TestWeatherForecastFilter(t *testing.T) {
	reported := time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)
	forecast := []WeatherForecast{{
		id:          "1",
		object_type: "forecast",
		areacode:    "123",
		title:       "Test Forecast",
		status:      "active",
		detail: Forecast{
			weather_today:    "Sunny",
			weather_tommorow: "Cloudy",
			max_temp:         30,
			min_temp:         20,
		},
		reported_at:    reported,
		info_domain:    "test_domain",
		info_object_id: "test_object_id",
	}}

	forecastlist := WeatherForecastlist{data: forecast}

	startTime, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2023-01-02T23:59:59Z")

	filterOptions := FilterOptions{
		StartTime: optional.Some(startTime),
		EndTime:   optional.Some(endTime),
	}

	filtered := forecastlist.Filter(filterOptions)

	if len(filtered.data) != 1 {
		t.Errorf("Expected forecast items to be '1', got '%d'", len(filtered.data))
	}
}

func TestWeatherEarthquakeFilter(t *testing.T) {
	earthquakes := []WeatherEarthquake{
		{reported_at: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)},
		{reported_at: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC)},
		{reported_at: time.Date(2023, 1, 3, 10, 0, 0, 0, time.UTC)},
	}

	list := WeatherEarthquakelist{data: earthquakes}

	startTime, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2023-01-02T23:59:59Z")

	filterOptions := FilterOptions{
		StartTime: optional.Some(startTime),
		EndTime:   optional.Some(endTime),
	}

	filtered := list.Filter(filterOptions)

	if len(filtered.data) != 2 {
		t.Errorf("Expected 3 earthquakes, got %d", len(filtered.data))
	}
}
