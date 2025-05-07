package fields

import (
	"context"

	"github.com/VeneLooool/fields-api/internal/model"
	"github.com/VeneLooool/fields-api/internal/pkg/db"
	"github.com/VeneLooool/fields-api/internal/pkg/ql"
	common "github.com/VeneLooool/fields-api/internal/repository"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/huandu/go-sqlbuilder"
)

type Repo struct {
	db db.DataBase
}

func New(db db.DataBase) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, field model.Field) (model.Field, error) {
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder().InsertInto(Table).Cols(
		Name.Short(),
		Culture.Short(),
		CreatedBy.Short(),
		Coordinates.Short(),
	).Values(
		field.Name,
		field.Culture,
		field.CreatedBy,
		field.Coordinates,
	)
	query, args := common.ReturningAll(ib).Build()

	var result model.Field
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Field{}, err
	}
	return result, nil
}

func (r *Repo) Update(ctx context.Context, field model.Field) (model.Field, error) {
	ub := sqlbuilder.PostgreSQL.NewUpdateBuilder().Update(Table)
	ub = ub.Set(ql.Fields{Name, Culture, Coordinates}.ToAssignments(ub, field.Name, field.Culture, field.Coordinates)...)
	ub = ub.Where(ub.Equal(ID.Short(), field.ID))
	query, args := common.ReturningAll(ub).Build()

	var result model.Field
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Field{}, err
	}
	return result, nil
}

func (r *Repo) Get(ctx context.Context, id uint64) (model.Field, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select(common.All()).From(Table)
	sb = sb.Where(sb.Equal(ID.Short(), id))
	query, args := sb.Build()

	var result model.Field
	if err := pgxscan.Get(ctx, r.db, &result, query, args...); err != nil {
		return model.Field{}, err
	}
	return result, nil
}

func (r *Repo) GetByAuthor(ctx context.Context, authorLogin string) ([]model.Field, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder().Select(common.All()).From(Table)
	sb = sb.Where(sb.Equal(CreatedBy.Short(), authorLogin))
	query, args := sb.Build()

	var result []model.Field
	if err := pgxscan.Select(ctx, r.db, &result, query, args...); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repo) Delete(ctx context.Context, id uint64) error {
	db := sqlbuilder.PostgreSQL.NewDeleteBuilder().DeleteFrom(Table)
	query, args := db.Where(db.Equal(ID.Short(), id)).Build()

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return err
	}
	return nil
}
