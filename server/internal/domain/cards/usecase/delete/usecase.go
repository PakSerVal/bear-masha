package delete

import (
	"context"

	"github.com/pkg/errors"
)

type Repository interface {
	Delete(ctx context.Context, id int64) error
}

type Usecase interface {
	Do(ctx context.Context, id int64) error
}

type usecase struct {
	repo Repository
}

func New(repo Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Do(ctx context.Context, id int64) error {
	err := u.repo.Delete(ctx, id)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
