package fields

import (
	"github.com/VeneLooool/fields-api/internal/model"
	proto_model "github.com/VeneLooool/fields-api/internal/pb/api/v1/model"
)

func transformFieldToModel(field *proto_model.Field) model.Field {
	return model.Field{
		ID:      field.GetId(),
		Name:    field.GetName(),
		Culture: field.GetCulture(),
	}
}

func transformFieldToProto(field model.Field) *proto_model.Field {
	return &proto_model.Field{
		Id:      field.ID,
		Name:    field.Name,
		Culture: field.Culture,
	}
}
