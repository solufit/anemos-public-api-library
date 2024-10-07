package libanemos

import (
	"testing"
	"time"

	"github.com/moznion/go-optional"
)

func TestWeatherWarningFilter(t *testing.T) {
	warnings := []WeatherWarning{
		{reported_at: "2023-01-01T10:00:00Z"},
		{reported_at: "2023-01-02T10:00:00Z"},
		{reported_at: "2023-01-03T10:00:00Z"},
	}

	list := WeatherWarninglist{data: warnings}

	startTime, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2023-01-02T23:59:59Z")

	filterOptions := FilterOptions{
		StartTime: optional.Some(startTime),
		EndTime:   optional.Some(endTime),
	}

	filtered := list.WeatherWarningFilter(filterOptions)

	if len(filtered.data) != 2 {
		t.Errorf("Expected 2 warnings, got %d", len(filtered.data))
	}
}

func (f WeatherForecast) WeatherForecastFilter(options FilterOptions) WeatherForecast {
	// Implement the filtering logic here
	return f
}

func TestWeatherForecastFilter(t *testing.T) {
	forecast := WeatherForecast{
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
		reported_at:    "2023-01-01T10:00:00Z",
		info_domain:    "test_domain",
		info_object_id: "test_object_id",
	}

	startTime, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2023-01-02T23:59:59Z")

	filterOptions := FilterOptions{
		StartTime: optional.Some(startTime),
		EndTime:   optional.Some(endTime),
	}

	filtered := forecast.WeatherForecastFilter(filterOptions)

	if filtered.id != "1" {
		t.Errorf("Expected forecast id to be '1', got '%s'", filtered.id)
	}
}

func TestWeatherEarthquakeFilter(t *testing.T) {
	earthquakes := []WeatherEarthquake{
		{reported_at: "2023-01-01T10:00:00Z"},
		{reported_at: "2023-01-02T10:00:00Z"},
		{reported_at: "2023-01-03T10:00:00Z"},
	}

	list := WeatherEarthquakelist{data: earthquakes}

	startTime, _ := time.Parse(time.RFC3339, "2023-01-01T00:00:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2023-01-02T23:59:59Z")

	filterOptions := FilterOptions{
		StartTime: optional.Some(startTime),
		EndTime:   optional.Some(endTime),
	}

	filtered := list.WeatherEarthquakeFilter(filterOptions)

	if len(filtered.data) != 3 {
		t.Errorf("Expected 3 earthquakes, got %d", len(filtered.data))
	}
}
