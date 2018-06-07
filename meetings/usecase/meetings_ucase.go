package usecase

import (
	"context"
	"time"
	"gopkg.in/mgo.v2/bson"

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

func (a *meetingUsecase) AddMeetingroom(c context.Context, m *models.MeetingRoom) (bson.ObjectId, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	id, err := a.meetingRepos.AddMeetingroom(ctx, m)
	if err != nil {
		return id, err
	}
	return id, nil
}
