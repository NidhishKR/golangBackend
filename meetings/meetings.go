package meetings

import (
	"CleanArchMeetingRoom/models"
	"gopkg.in/mgo.v2/bson"
	"context"
)

type MeetingsRepository interface {
	GetByRegion(ctx context.Context, id string) (*[]models.MeetingRoom, error)
	AddMeetingroom(ctx context.Context, a *models.MeetingRoom) (bson.ObjectId, error)
}
