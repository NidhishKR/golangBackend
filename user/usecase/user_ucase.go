package usecase

import (
	"context"
	"time"
	"CleanArchMeetingRoom/user"
)

type meetingUsecase struct {
	userRepos   user.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u user.UserRepository, timeout time.Duration) user.UserUsecase {
	return &meetingUsecase{
		userRepos:   u,
		contextTimeout: timeout,
	}
}

func (a *meetingUsecase) GetUserByAuthToken(c context.Context, token string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err := a.userRepos.GetUserByAuthToken(ctx, token)
	if err != nil {
		return false, err
	}
	return res, nil
}