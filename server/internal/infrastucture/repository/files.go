package repository

import (
	"context"

	"github.com/PakSerVal/bear-masha/internal/domain/file"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type fileRepository struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) *fileRepository {
	return &fileRepository{
		db: db,
	}
}

func (f *fileRepository) Save(ctx context.Context, path string, fileType string) (*file.File, error) {
	query := `insert into files(path, type) values ($1, $2) returning id`

	var id int64
	err := f.db.QueryRowContext(ctx, query, path, fileType).Scan(&id)
	if err != nil {
		return nil, errors.Wrap(err, "repo: create card")
	}

	return &file.File{
		ID:   id,
		Path: path,
		Type: fileType,
	}, nil
}

func (f *fileRepository) GetCardFiles(ctx context.Context, cardId int64) ([]file.File, error) {
	query := `select f.id, f.path, f.type from files f
    inner join card_files cf on f.id = cf.file_id
	where cf.card_id = $1`

	var files []file.File
	err := f.db.SelectContext(ctx, &files, query, cardId)
	if err != nil {
		return nil, errors.Wrap(err, "repo: get card files")
	}

	return files, nil
}
