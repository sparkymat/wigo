package route

import (
	"context"

	"github.com/sparkymat/wigo/internal/dbx"
)

type UserService interface {
	FetchUserByUsername(ctx context.Context, username string) (dbx.User, error)
	CreateUser(ctx context.Context, name string, username string, encryptedPassword string) (dbx.User, error)
}
