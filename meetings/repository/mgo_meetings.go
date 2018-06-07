package repository

import (
	"context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
	meeting "CleanArchMeetingRoom/meetings"
	models "CleanArchMeetingRoom/models"
	config "CleanArchMeetingRoom/config"
)

const MEETINGROOM string = "MEETINGROOM"

type mgoMeetingsRepository struct {
	Conn *mgo.Database
}

func NewMgoMeetingsRepository(Conn *mgo.Database) meeting.MeetingsRepository {
	return &mgoMeetingsRepository{Conn}
}

func (m *mgoMeetingsRepository) GetByRegion(ctx context.Context, region string) (*[]models.MeetingRoom, error) {
	cn := config.CollectionNames(MEETINGROOM)
	c := m.Conn.C(cn)
	var meetingRoom []models.MeetingRoom
	iter := c.Find(bson.M{"region": region}).Limit(100).Iter()
	err := iter.All(&meetingRoom)
	if err != nil {
		return  &[]models.MeetingRoom{}, models.NOT_FOUND_ERROR
	}
	return &meetingRoom, nil
}