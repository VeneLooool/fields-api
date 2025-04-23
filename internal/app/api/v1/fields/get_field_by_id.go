package fields

import (
	"context"
	"errors"

	desc "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
	"github.com/VeneLooool/fields-api/internal/pkg/error_hub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetFieldByID(ctx context.Context, req *desc.GetFieldByID_Request) (*desc.GetFieldByID_Response, error) {
	field, err := i.fieldUC.Get(ctx, req.GetId())
	if err != nil {
		code := codes.Internal
		if errors.Is(err, error_hub.ErrFieldNotFound) {
			code = codes.NotFound
		}
		return nil, status.Error(code, err.Error())
	}
	return &desc.GetFieldByID_Response{
		Field: transformFieldToProto(field),
	}, nil
}
