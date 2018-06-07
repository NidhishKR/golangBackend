package user

import "context"

type UserUsecase interface {
	GetUserByAuthToken(ctx context.Context, token string) (bool, error)
}
