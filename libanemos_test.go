package libanemos

import (
	"testing"
)

type mockAnemosData struct{}

func (m mockAnemosData) Filter() anemosData[any] {
	return m
}

// Removed the conflicting AnemosGet struct declaration

// Removed the duplicate NewAnemosGet function declaration

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
