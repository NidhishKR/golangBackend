package meetings

import (
	"CleanArchMeetingRoom/models"
	"gopkg.in/mgo.v2/bson"
	"context"
)

type MeetingsRepository interface {
	GetByRegion(ctx context.Context, id string) (*[]models.MeetingRoom, error)
	AddMeetingroom(ctx context.Context, mr *models.MeetingRoom) (bson.ObjectId, error)
	AddMeeting(ctx context.Context, m *models.NewMeeting, u *models.GlobalUser) error
	GetConcurrentMeetings(m *models.NewMeeting) ([]models.NewMeeting, error)
}
