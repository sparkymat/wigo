package user

import (
	"context"
	"fmt"

	"github.com/sparkymat/wigo/internal/dbx"
	"github.com/sparkymat/wigo/internal/service"
)

func New(db service.DatabaseProvider) *Service {
	return &Service{
		db: db,
	}
}

type Service struct {
	db service.DatabaseProvider
}

func (s *Service) CreateUser(ctx context.Context, name string, username string, encryptedPassword string) (dbx.User, error) {
	user, err := s.db.CreateUser(ctx, dbx.CreateUserParams{
		Name:              name,
		Username:          username,
		EncryptedPassword: encryptedPassword,
	})
	if err != nil {
		return dbx.User{}, fmt.Errorf("unable to fetch user. err: %w", err)
	}

	return user, nil
}

func (s *Service) FetchUserByUsername(ctx context.Context, username string) (dbx.User, error) {
	user, err := s.db.FetchUserByUsername(ctx, username)
	if err != nil {
		return user, fmt.Errorf("unable to fetch user. err: %w", err)
	}

	return user, nil
}
