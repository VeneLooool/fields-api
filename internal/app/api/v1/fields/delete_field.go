package fields

import (
	"context"
	
	desc "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteField(ctx context.Context, req *desc.DeleteField_Request) (*emptypb.Empty, error) {
	if err := i.fieldUC.Delete(ctx, req.GetId()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
