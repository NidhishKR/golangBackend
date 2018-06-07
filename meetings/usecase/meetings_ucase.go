package usecase

import (
	"context"
	"time"
	
	"CleanArchMeetingRoom/models"
	"CleanArchMeetingRoom/meetings"
)

type meetingUsecase struct {
	meetingRepos   meetings.MeetingsRepository
	contextTimeout time.Duration
}

func NewMeetingUsecase(m meetings.MeetingsRepository, timeout time.Duration) meetings.MeetingsUsecase {
	return &meetingUsecase{
		meetingRepos:   m,
		contextTimeout: timeout,
	}
}

func (a *meetingUsecase) GetByRegion(c context.Context, region string) (*[]models.MeetingRoom, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err := a.meetingRepos.GetByRegion(ctx, region)
	if err != nil {
		return nil, err
	}
	return res, nil
}