package libanemos

type AnemosGet struct {
	Data []anemosData[any]
}

type anemosData[T any] interface {
	Filter() anemosData[T]
}

func NewAnemosGet() *AnemosGet {
	anemosget := AnemosGet{
		Data: []anemosData[any]{},
	}
	return &anemosget
}
