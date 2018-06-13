package usecase

import (
	"context"
	"time"
	"gopkg.in/mgo.v2/bson"

	"CleanArchMeetingRoom/models"
	"CleanArchMeetingRoom/meetings"
	"CleanArchMeetingRoom/user"

	// "fmt"

)

type meetingUsecase struct {
	meetingRepos   meetings.MeetingsRepository
	userRepos user.UserRepository
	contextTimeout time.Duration
}

func NewMeetingUsecase(m meetings.MeetingsRepository, u user.UserRepository, timeout time.Duration) meetings.MeetingsUsecase {
	return &meetingUsecase{
		meetingRepos:   m,
		userRepos: u,
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

func (a *meetingUsecase) AddMeeting(c context.Context, m *models.NewMeeting)  (*models.NewMeeting, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	user, err := a.userRepos.GetUserByEmailId(c, m.BookedBy)
	if err != nil {
		return nil, err
	}
	id, err := a.meetingRepos.AddMeeting(ctx, m, &user)
	if err != nil {
		return nil, err
	}
	m.Id = id
	return m, nil
}
