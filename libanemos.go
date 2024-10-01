// Package libanemos provides structures and functions to work with Anemos data.
package libanemos

// AnemosGet is a structure that holds a slice of anemosData of any type.
type AnemosGet struct {
	Data []anemosData[any]
}

// The anemosData interface defines a generic type T and requires a Filter method
// that returns an anemosData of the same type.
type anemosData[T any] interface {
	Filter() anemosData[T]
}

// NewAnemosGet initializes and returns a pointer to an AnemosGet instance with an
// empty slice of anemosData.
func NewAnemosGet() *AnemosGet {
	anemosget := AnemosGet{
		Data: []anemosData[any]{},
	}
	return &anemosget
}
