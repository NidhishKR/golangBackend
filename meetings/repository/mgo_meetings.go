package repository

import (
	"context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
	meeting "CleanArchMeetingRoom/meetings"
	models "CleanArchMeetingRoom/models"
)

type mgoMeetingsRepository struct {
	Conn *mgo.Database
}

func NewMgoMeetingsRepository(Conn *mgo.Database) meeting.MeetingsRepository {
	return &mgoMeetingsRepository{Conn}
}

func (m *mgoMeetingsRepository) GetByRegion(ctx context.Context, region string) (*[]models.MeetingRoom, error) {
	cn := models.MEETINGROOM.DB.MODELS.COLLECTION
	c := m.Conn.C(cn)
	var meetingRoom []models.MeetingRoom
	iter := c.Find(bson.M{"region": region}).Limit(100).Iter()
	err := iter.All(&meetingRoom)
	if err != nil {
		return  &[]models.MeetingRoom{}, models.NOT_FOUND_ERROR
	}
	return &meetingRoom, nil
}

func (m *mgoMeetingsRepository) AddMeetingroom(ctx context.Context, mmr *models.MeetingRoom) (bson.ObjectId, error) {
	cn := models.MEETINGROOM.DB.MODELS.COLLECTION
	c := m.Conn.C(cn)
	id := bson.NewObjectId()
	mr := models.MeetingRoom{
		Id : id,
		Name:mmr.Name,
		Region :mmr.Region,		
		Building :mmr.Building,
		CreatedBy :mmr.CreatedBy,
		CreatedAt :mmr.CreatedAt,
	}
	err := c.Insert(mr)

	if err != nil {
		return id, err
	}
	return id, nil
}

func (m *mgoMeetingsRepository) AddMeeting(ctx context.Context, mm *models.NewMeeting) (bson.ObjectId, error) {
	cn := models.MEETINGROOM.DB.MODELS.COLLECTION
	c := m.Conn.C(cn)
	id := bson.NewObjectId()

	nm := models.NewMeeting{
		MeetingSubject : mm.MeetingSubject,
		Meetingroom : mm.Meetingroom,		
		MeetingStartTime : mm.MeetingStartTime,
		MeetingEndTime : mm.MeetingEndTime,
		MeetingDate : mm.MeetingDate,
		BookedBy : mm.BookedBy,
		MeetingStatus : mm.MeetingStatus,
		MeetingReccurance : mm.MeetingReccurance,
	}

	err := c.Insert(nm)

	if err != nil {
		return id, err
	}
	return id, nil
}