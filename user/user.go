package user

import (
	"context"

	model "CleanArchMeetingRoom/models"
)

type UserRepository interface {
	GetUserByAuthToken(ctx context.Context, token string) (bool, error)
	GetUserByEmailId(ctx context.Context, emailId string) (model.GlobalUser, error)
}
