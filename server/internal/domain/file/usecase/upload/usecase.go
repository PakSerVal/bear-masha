package upload

import (
	"context"

	"github.com/PakSerVal/bear-masha/internal/domain/file"
	"github.com/pkg/errors"
)

type Repository interface {
	Save(ctx context.Context, path string, fileType string) (*file.File, error)
}

type Usecase interface {
	Do(ctx context.Context, path string, fileType string) (*file.File, error)
}

type usecase struct {
	repo Repository
}

func New(repo Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Do(ctx context.Context, path string, fileType string) (*file.File, error) {
	if fileType != file.TypeImageMain && fileType != file.TypeImage {
		return nil, errors.New("invalid file type")
	}

	f, err := u.repo.Save(ctx, path, fileType)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return f, nil
}
