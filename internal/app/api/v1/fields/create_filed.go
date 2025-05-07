package fields

import (
	"context"

	desc "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateField(ctx context.Context, req *desc.CreateField_Request) (*desc.CreateField_Response, error) {
	field, err := i.fieldUC.Create(ctx, transformCreateFieldReqToModel(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateField_Response{
		Field: transformFieldToProto(field),
	}, nil
}
