package user

import "context"

type UserRepository interface {
	GetUserByAuthToken(ctx context.Context, token string) (bool, error)
}
