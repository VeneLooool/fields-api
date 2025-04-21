package fields

import (
	desc "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
)

// Implementation is a Service implementation
type Implementation struct {
	desc.UnimplementedFieldsServer
}

// NewService return new instance of Implementation.
func NewService() *Implementation {
	return &Implementation{}
}
