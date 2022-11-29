package create_test

import (
	"context"
	"testing"
	"time"

	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/create"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/create/mocks"
	"github.com/pkg/errors"
	mock2 "github.com/stretchr/testify/mock"
)

func Test_usecase_Do(t *testing.T) {
	type fields struct {
		cardRepo      create.CardRepository
		cardFilesRepo create.CardFilesRepository
	}
	type args struct {
		ctx         context.Context
		title       string
		description string
		deadlineAt  time.Time
		files       *[]int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Card repo error",
			fields: fields{
				cardRepo: func() create.CardRepository {
					mock := &mocks.CardRepository{}
					mock.EXPECT().Create(mock2.Anything, mock2.Anything).Return(nil, errors.New("some error"))

					return mock
				}(),
				cardFilesRepo: nil,
			},
			args: args{
				ctx:         context.TODO(),
				title:       "title",
				description: "description",
				deadlineAt:  time.Now(),
				files:       nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := create.New(tt.fields.cardRepo, tt.fields.cardFilesRepo)
			if err := u.Do(tt.args.ctx, tt.args.title, tt.args.description, tt.args.deadlineAt, tt.args.files); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
