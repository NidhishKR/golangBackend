package usecase

import (
	"context"
	"time"
	"gopkg.in/mgo.v2/bson"
	"fmt"

	"CleanArchMeetingRoom/models"
	"CleanArchMeetingRoom/meetings"
	"CleanArchMeetingRoom/user"
	"CleanArchMeetingRoom/utils"

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
	vmd, err := utils.IsValidTime(m.MeetingDate)
	vst, err := utils.IsValidTime(m.MeetingStartTime)
	vet, err := utils.IsValidTime(m.MeetingEndTime) 
	if err != nil {
		return nil, models.INVALID_TIME_INPUT
	}
	recurrence := models.ReccuranceCount{}
	recurrence.Id = fmt.Sprintf("%d", time.Now().Unix())
	recurrence.Repetition = m.MeetingReccurance
	for n := 0; n <= m.MeetingReccurance; n++ {
		m.Id = fmt.Sprintf("%d", time.Now().UnixNano())
		m.MeetingStartTime = vst.Add(time.Hour * 24 * time.Duration(n)).Format(time.RFC3339)
		m.MeetingEndTime = vet.Add(time.Hour * 24 * time.Duration(n)).Format(time.RFC3339)
		m.MeetingDate = vmd.Add(time.Hour * 24 * time.Duration(n)).Format(time.RFC3339)
		m.ReccuranceCount = recurrence
		if existing, err := checkTiming(a, m, false); (existing && err !=nil) {
			return nil, models.TIME_SLOT_BOOKED
		}
		res := a.meetingRepos.AddMeeting(ctx, m, &user)
		if res != nil {
			return nil, res
		}
	}
	return m, nil
}
