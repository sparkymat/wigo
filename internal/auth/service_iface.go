package auth

import (
	"context"

	"github.com/sparkymat/wigo/internal/dbx"
)

type ConfigService interface {
	JWTSecret() string
	ProxyAuthUsernameHeader() string
	ProxyAuthNameHeader() string
}

type UserService interface {
	FetchUserByUsername(ctx context.Context, username string) (dbx.User, error)
	CreateUser(ctx context.Context, name string, username string, encryptedPassword string) (dbx.User, error)
}
