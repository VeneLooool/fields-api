package fields

import (
	desc "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
)

// Implementation is a Service implementation
type Implementation struct {
	desc.UnimplementedFieldsServer

	fieldUC FieldUC
}

// NewService return new instance of Implementation.
func NewService(fieldUC FieldUC) *Implementation {
	return &Implementation{
		fieldUC: fieldUC,
	}
}
