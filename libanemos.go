package libanemos

type AnemosGet struct {
	Data []anemosData[any]
}

type anemosData[T any] interface {
	Filter() anemosData[T]
}

func NewAnemosGet() *AnemosGet {
	return &AnemosGet{
		Data: []anemosData[any]{
			// Add Anemos Data Type Here
		},
	}
}
