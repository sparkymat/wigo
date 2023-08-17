package service

import (
	"context"

	"github.com/sparkymat/wigo/internal/dbx"
)

//nolint:interfacebloat
type DatabaseProvider interface {
	CreateUser(ctx context.Context, arg dbx.CreateUserParams) (dbx.User, error)
	FetchUserByUsername(ctx context.Context, username string) (dbx.User, error)
}
