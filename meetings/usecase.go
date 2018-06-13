package meetings

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	model "CleanArchMeetingRoom/models"
)

type MeetingsUsecase interface {
	GetByRegion(ctx context.Context, region string) (*[]model.MeetingRoom, error)
	AddMeetingroom(context.Context, *model.MeetingRoom) (bson.ObjectId, error)
	AddMeeting(context.Context, *model.NewMeeting)  (*model.NewMeeting, error)
}