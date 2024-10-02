// Package libanemos provides structures and functions to work with Anemos data.
package libanemos

import (
	"time"

	"github.com/moznion/go-optional"
)

// AnemosGet is a structure that holds a slice of anemosData of any type.
type AnemosGet struct {
	Data []anemosData[any]
}

// The anemosData interface defines a generic type T and requires a Filter method
// that returns an anemosData of the same type.
type anemosData[T any] interface {
	Filter(FilterOptions) anemosData[T]
}

// NewAnemosGet initializes and returns a pointer to an AnemosGet instance with an
// empty slice of anemosData.
func NewAnemosGet() *AnemosGet {
	anemosget := AnemosGet{
		Data: []anemosData[any]{},
	}
	return &anemosget
}

// PostCode is a type alias for a string representing a postal code.
type PostCode string

// FilterOptions is a structure that holds optional values for filtering anemosData.
type FilterOptions struct {
	PostCode  optional.Option[PostCode]
	StartTime optional.Option[time.Time]
	EndTime   optional.Option[time.Time]
}
