package get_by_id

import (
	"context"

	"github.com/PakSerVal/bear-masha/internal/domain/cards"
	"github.com/PakSerVal/bear-masha/internal/domain/file"
	"github.com/pkg/errors"
)

type CardRepository interface {
	GetByID(ctx context.Context, id int64) (*cards.Card, error)
}

type FileRepository interface {
	GetCardFiles(ctx context.Context, cardID int64) ([]file.File, error)
}

type Usecase interface {
	Do(ctx context.Context, id int64) (*cards.Card, error)
}

type usecase struct {
	cardRepo CardRepository
	fileRepo FileRepository
}

func New(cardRepo CardRepository, fileRepo FileRepository) Usecase {
	return &usecase{
		cardRepo: cardRepo,
		fileRepo: fileRepo,
	}
}

func (u *usecase) Do(ctx context.Context, id int64) (*cards.Card, error) {
	card, err := u.cardRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if card == nil {
		return nil, nil
	}

	card.Files, err = u.fileRepo.GetCardFiles(ctx, card.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return card, nil
}
