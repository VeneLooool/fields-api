package fields

import (
	"context"
	"errors"
	"github.com/VeneLooool/fields-api/internal/model"
	"github.com/VeneLooool/fields-api/internal/pkg/error_hub"
	"github.com/jackc/pgx/v4"
)

type UseCase struct {
	repo Repo
}

func New(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) Create(ctx context.Context, field model.Field) (model.Field, error) {
	return u.repo.Create(ctx, field)
}

func (u *UseCase) Update(ctx context.Context, field model.Field) (model.Field, error) {
	return u.repo.Update(ctx, field)
}

func (u *UseCase) Get(ctx context.Context, id uint64) (model.Field, error) {
	field, err := u.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Field{}, error_hub.ErrFieldNotFound
		}
		return model.Field{}, err
	}
	return field, nil
}

func (u *UseCase) GetByAuthor(ctx context.Context, authorLogin string) ([]model.Field, error) {
	return u.repo.GetByAuthor(ctx, authorLogin)
}

func (u *UseCase) Delete(ctx context.Context, id uint64) error {
	return u.repo.Delete(ctx, id)
}
