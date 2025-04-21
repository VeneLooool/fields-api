package fields

import (
	"context"
	desc "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
	desc_model "github.com/VeneLooool/fields-api/internal/pb/api/v1/model"
)

func (i *Implementation) GetFieldByID(ctx context.Context, req *desc.GetFieldByID_Request) (*desc.GetFieldByID_Response, error) {

	return &desc.GetFieldByID_Response{
		Field: &desc_model.Field{
			Id:      1,
			Name:    "test",
			Culture: "test",
		},
	}, nil
}
