package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type cardFilesRepository struct {
	db *sqlx.DB
}

func NewCardFileRepository(db *sqlx.DB) *cardFilesRepository {
	return &cardFilesRepository{
		db: db,
	}
}

func (f *cardFilesRepository) Save(ctx context.Context, cardId int64, fileId int64) error {
	query := `insert into card_files(card_id, file_id) values ($1, $2)`

	_, err := f.db.ExecContext(ctx, query, cardId, fileId)
	if err != nil {
		return errors.Wrap(err, "repo: create card")
	}

	return nil
}
