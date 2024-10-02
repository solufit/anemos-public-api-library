package libanemos

import (
	"testing"
	"time"

	"github.com/moznion/go-optional"
)

type mockAnemosData struct{}

func (m mockAnemosData) Filter(filterOption FilterOptions) anemosData[any] {
	return m
}

func TestNewAnemosGet(t *testing.T) {
	anemosGet := NewAnemosGet()

	if anemosGet == nil {
		t.Fatalf("Expected non-nil AnemosGet, got nil")
	}
	anemosGet.Data = append(anemosGet.Data, mockAnemosData{})

	// Check anemosGet.Data Type
	if len(anemosGet.Data) == 0 {
		t.Fatalf("Expected Data length not to be 0")
	}

	if _, ok := anemosGet.Data[0].(mockAnemosData); !ok {
		t.Fatalf("Expected Data[0] to be of type mockAnemosData")
	}
}

func TestFilterOptions(t *testing.T) {
	postCode := PostCode("12345")
	startTime := time.Now().Add(-time.Hour)
	endTime := time.Now()

	filterOptions := FilterOptions{
		PostCode:  optional.Some(postCode),
		StartTime: optional.Some(startTime),
		EndTime:   optional.Some(endTime),
	}

	postcode_value, err := filterOptions.PostCode.Take()
	if err != nil {
		t.Fatalf("Expected PostCode to be set, got error: %v", err)
	}
	if postcode_value != postCode {
		t.Fatalf("Expected PostCode to be %v, got %v", postCode, postcode_value)
	}

	starttime_value, err := filterOptions.StartTime.Take()
	if err != nil {
		t.Fatalf("Expected StartTime to be set, got error: %v", err)
	}
	if starttime_value != startTime {
		t.Fatalf("Expected StartTime to be %v, got %v", startTime, starttime_value)
	}

	endtime_value, err := filterOptions.EndTime.Take()
	if err != nil {
		t.Fatalf("Expected EndTime to be set, got error: %v", err)
	}

	if endtime_value != endTime {
		t.Fatalf("Expected EndTime to be %v, got %v", endTime, endtime_value)
	}
}
