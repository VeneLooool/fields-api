package fields

import (
	"context"

	"github.com/VeneLooool/fields-api/internal/model"
	desc "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) UpdateField(ctx context.Context, req *desc.UpdateField_Request) (*desc.UpdateField_Response, error) {
	field, err := i.fieldUC.Update(ctx, model.Field{
		ID:      req.GetId(),
		Name:    req.GetName(),
		Culture: req.GetCulture(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.UpdateField_Response{
		Field: transformFieldToProto(field),
	}, nil
}
