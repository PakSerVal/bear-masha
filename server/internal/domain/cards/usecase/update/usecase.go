package update

import (
	"context"
	"time"

	"github.com/PakSerVal/bear-masha/internal/domain/cards"
	"github.com/pkg/errors"
)

type Repository interface {
	Update(ctx context.Context, card cards.Card) error
	GetByID(ctx context.Context, id int64) (*cards.Card, error)
}

type Usecase interface {
	Do(ctx context.Context, id int64, title *string, status *string, description *string, deadlineAt *time.Time) error
}

type usecase struct {
	repo Repository
}

func New(repo Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Do(ctx context.Context, id int64, title *string, status *string, description *string, deadlineAt *time.Time) error {
	card, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return errors.WithStack(err)
	}

	if card == nil {
		return errors.New("card not found")
	}

	if err != nil {
		return errors.WithStack(err)
	}

	if title != nil {
		card.Title = *title
	}

	if status != nil {
		card.Status = *status
	}

	if description != nil {
		card.Description = *description
	}

	if deadlineAt != nil {
		card.DeadlineAt = *deadlineAt
	}

	err = u.repo.Update(ctx, *card)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
