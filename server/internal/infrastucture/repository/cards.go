package repository

import (
	"context"
	"database/sql"

	"github.com/PakSerVal/bear-masha/internal/domain/cards"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type cardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *cardRepository {
	return &cardRepository{
		db: db,
	}
}

func (c *cardRepository) GetList(ctx context.Context, limit int64) ([]cards.Card, error) {
	const query = `
		select
			id,
			title,
			description,
			status,
			created_at, 
			deadline_at
 		from cards
		order by created_at 
		limit $1
	`

	var cards []cards.Card

	err := c.db.SelectContext(ctx, &cards, query, limit)
	if err != nil {
		return nil, errors.Wrap(err, "repo: get list")
	}

	return cards, nil
}

func (c *cardRepository) GetByID(ctx context.Context, id int64) (*cards.Card, error) {
	const query = `
		select
			id,
			title,
			description,
			status,
			created_at, 
			deadline_at
 		from cards
		where id = $1
	`

	var card cards.Card
	err := c.db.GetContext(ctx, &card, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "repo: get card by id")
	}

	return &card, nil
}

func (c *cardRepository) Create(ctx context.Context, card cards.Card) (*cards.Card, error) {
	query := `
		insert into cards(
		    title,
		    description,
		    status,
		    deadline_at
		) values ($1, $2, $3, $4) 
		returning id`

	var id int64

	err := c.db.QueryRowContext(ctx, query, card.Title, card.Description, card.Status, card.DeadlineAt).Scan(&id)
	if err != nil {
		return nil, errors.Wrap(err, "repo: create card")
	}

	card.ID = id

	return &card, nil
}

func (c *cardRepository) Delete(ctx context.Context, id int64) error {
	query := `delete from cards where id = $1`

	_, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "repo: delete card")
	}

	return nil
}

func (c *cardRepository) Update(ctx context.Context, card cards.Card) error {
	query := `UPDATE cards SET title=$1, description=$2, status=$3, deadline_at=$4 WHERE id=$5`

	_, err := c.db.ExecContext(ctx, query, card.Title, card.Description, card.Status, card.DeadlineAt, card.ID)
	if err != nil {
		return errors.Wrap(err, "repo: update card")
	}

	return nil
}
