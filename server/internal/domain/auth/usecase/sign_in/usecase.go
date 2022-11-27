package sign_in

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type Usecase interface {
	Do(ctx context.Context, login string, password string) (*string, error)
}

type usecase struct {
	login     string
	password  string
	secretKey string
}

func New(login string, password string, secretKey string) Usecase {
	return &usecase{
		login:     login,
		password:  password,
		secretKey: secretKey,
	}
}

func (u *usecase) Do(ctx context.Context, login string, password string) (*string, error) {
	if u.login != login || u.password != password {
		return nil, errors.New("invalid loging or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.secretKey))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &tokenString, nil
}
