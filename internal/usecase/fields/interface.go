package fields

import (
	"context"

	"github.com/VeneLooool/fields-api/internal/model"
)

type Repo interface {
	Create(ctx context.Context, field model.Field) (model.Field, error)
	Update(ctx context.Context, field model.Field) (model.Field, error)
	Get(ctx context.Context, id uint64) (model.Field, error)
	Delete(ctx context.Context, id uint64) error
}
