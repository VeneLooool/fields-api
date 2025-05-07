package fields

import (
	"github.com/VeneLooool/fields-api/internal/model"
	desc "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
	proto_model "github.com/VeneLooool/fields-api/internal/pb/api/v1/model"
)

func transformFieldToModel(field *proto_model.Field) model.Field {
	return model.Field{
		ID:        field.GetId(),
		Name:      field.GetName(),
		Culture:   field.GetCulture(),
		CreatedBy: field.GetCreatedBy(),
	}
}

func transformCreateFieldReqToModel(req *desc.CreateField_Request) model.Field {
	return model.Field{
		Name:        req.GetName(),
		Culture:     req.GetCulture(),
		CreatedBy:   req.GetCreatedBy(),
		Coordinates: transformCoordinatesToModel(req.GetCoordinates()),
	}
}

func transformUpdateFieldReqToModel(req *desc.UpdateField_Request) model.Field {
	return model.Field{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Culture:     req.GetCulture(),
		Coordinates: transformCoordinatesToModel(req.GetCoordinates()),
	}
}

func transformCoordinatesToModel(protoCoordinates []*proto_model.Coordinate) []model.Coordinate {
	coordinates := make([]model.Coordinate, 0, len(protoCoordinates))
	for _, coordinate := range protoCoordinates {
		coordinates = append(coordinates, model.Coordinate{
			Latitude:  coordinate.GetLatitude(),
			Longitude: coordinate.GetLongitude(),
		})
	}
	return coordinates
}

func transformFieldsToProto(fields []model.Field) []*proto_model.Field {
	protoFields := make([]*proto_model.Field, 0, len(fields))
	for _, field := range fields {
		protoFields = append(protoFields, transformFieldToProto(field))
	}
	return protoFields
}

func transformFieldToProto(field model.Field) *proto_model.Field {
	return &proto_model.Field{
		Id:          field.ID,
		Name:        field.Name,
		Culture:     field.Culture,
		CreatedBy:   field.CreatedBy,
		Coordinates: transformCoordinatesToProto(field.Coordinates),
	}
}

func transformCoordinatesToProto(coordinates []model.Coordinate) []*proto_model.Coordinate {
	protoCoordinates := make([]*proto_model.Coordinate, 0, len(coordinates))
	for _, coordinate := range coordinates {
		protoCoordinates = append(protoCoordinates, &proto_model.Coordinate{
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		})
	}
	return protoCoordinates
}
