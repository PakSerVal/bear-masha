package get_list

import (
	"context"

	"github.com/PakSerVal/bear-masha/internal/domain/cards"
	"github.com/PakSerVal/bear-masha/internal/domain/file"
	"github.com/pkg/errors"
)

const cardLimit = 20

type CardRepository interface {
	GetList(ctx context.Context, limit int64) ([]cards.Card, error)
}

type FileRepository interface {
	GetCardFiles(ctx context.Context, cardId int64) ([]file.File, error)
}

type Usecase interface {
	Do(ctx context.Context) (map[string][]cards.Card, error)
}

type usecase struct {
	cardRepo CardRepository
	fileRepo FileRepository
}

func New(repo CardRepository, fileRepo FileRepository) *usecase {
	return &usecase{
		cardRepo: repo,
		fileRepo: fileRepo,
	}
}

func (u *usecase) Do(ctx context.Context) (map[string][]cards.Card, error) {
	res, err := u.cardRepo.GetList(ctx, cardLimit)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	cardByStatuses := map[string][]cards.Card{}
	for _, card := range res {
		card.Files, err = u.fileRepo.GetCardFiles(ctx, card.Id)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		cardByStatuses[card.Status] = append(cardByStatuses[card.Status], card)
	}

	return cardByStatuses, nil
}
