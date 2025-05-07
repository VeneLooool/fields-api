package fields

import (
	"context"
	
	desc "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetFieldsByAuthor(ctx context.Context, req *desc.GetFieldsByAuthor_Request) (*desc.GetFieldsByAuthor_Response, error) {
	fields, err := i.fieldUC.GetByAuthor(ctx, req.GetLogin())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &desc.GetFieldsByAuthor_Response{
		Fields: transformFieldsToProto(fields),
	}, nil
}
