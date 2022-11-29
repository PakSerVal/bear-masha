package validate

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type Usecase interface {
	Do(ctx context.Context, tokenString string) error
}

type usecase struct {
	secretKey string
}

func New(secretKey string) Usecase {
	return &usecase{secretKey: secretKey}
}

func (u *usecase) Do(ctx context.Context, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(u.secretKey), nil
	})
	if err != nil {
		return errors.Wrap(err, "invalid token")
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
