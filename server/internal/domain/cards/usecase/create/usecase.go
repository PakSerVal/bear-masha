package create

import (
	"context"
	"time"

	"github.com/PakSerVal/bear-masha/internal/domain/cards"
	"github.com/pkg/errors"
)

//go:generate mockery --name=CardRepository --with-expecter
type CardRepository interface {
	Create(ctx context.Context, card cards.Card) (*cards.Card, error)
}

type CardFilesRepository interface {
	Save(ctx context.Context, cardId int64, fileId int64) error
}

type Usecase interface {
	Do(ctx context.Context, title string, description string, deadlineAt time.Time, files *[]int64) error
}

type usecase struct {
	cardRepo      CardRepository
	cardFilesRepo CardFilesRepository
}

func New(repo CardRepository, cardFilesRepo CardFilesRepository) *usecase {
	return &usecase{
		cardRepo:      repo,
		cardFilesRepo: cardFilesRepo,
	}
}

func (u *usecase) Do(ctx context.Context, title string, description string, deadlineAt time.Time, files *[]int64) error {
	card, err := u.cardRepo.Create(ctx, cards.Card{
		Title:       title,
		Description: description,
		Status:      cards.StatusTodo,
		DeadlineAt:  deadlineAt,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	if files != nil {
		for _, fileId := range *files {
			err := u.cardFilesRepo.Save(ctx, card.Id, fileId)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}
