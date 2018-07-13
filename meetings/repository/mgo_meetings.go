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

func (m *mgoMeetingsRepository) AddMeeting(ctx context.Context, mm *models.NewMeeting, u *models.GlobalUser) error {
	cn := models.MEETINGROOM.DB.MODELS.COLLECTION
	c := m.Conn.C(cn)
	um := models.GlobalUser{
		EmployeeName:  u.EmployeeName,
		UserId:  u.UserId,
		EmailId: u.EmailId,
		ProfileImageUrl: u.ProfileImageUrl,
	} 
	nm := models.NewMeeting{
		Id: mm.Id,
		MeetingSubject : mm.MeetingSubject,
		Meetingroom : mm.Meetingroom,		
		MeetingStartTime : mm.MeetingStartTime,
		MeetingEndTime : mm.MeetingEndTime,
		MeetingDate : mm.MeetingDate,
		BookedBy : mm.BookedBy,
		UserDetails: &um,
		MeetingStatus : mm.MeetingStatus,
		MeetingReccurance : mm.MeetingReccurance,
		ReccuranceCount : mm.ReccuranceCount
	}
	Who := bson.M{"_id": bson.ObjectIdHex(mm.Meetingroom)}
	PushToArray := bson.M{"$push": bson.M{"meetings": nm}}
	err := c.Update(Who, PushToArray)
	if err != nil {
		return models.INTERNAL_SERVER_ERROR
	}
	return nil
}

func (m *mgoMeetingsRepository) GetConcurrentMeetings(mm *models.NewMeeting) ([]models.NewMeeting, error){
	cn := models.MEETINGROOM.DB.MODELS.COLLECTION
	c := m.Conn.C(cn)
	pipeline := []bson.M{
		{"$unwind": "$meetings"},
		{"$match": bson.M{ "meetings.meetingroom" : mm.Meetingroom,"meetings.meetingDate" : mm.MeetingDate}},
		{"$group": bson.M{"_id": "$_id",
		 "meetings": bson.M{"$addToSet": "$meetings"}},
		},
	}
	rs := models.MeetingRes{}
	pipe := c.Pipe(pipeline)
	err := pipe.One(&rs)
	if(err != nil){
		return []models.NewMeeting{}, err
	}
	return rs.Meetings, nil
}

func (m *mgoMeetingsRepository) GetMeetingsByDateRange(ctx context.Context, startDate string, endDate string,  id string) (bson.M, error) {
	cn := models.MEETINGROOM.DB.MODELS.COLLECTION
	c := m.Conn.C(cn)
	pipeline := []bson.M{
		{"$unwind": "$meetings"},
		{"$match": bson.M{"_id": bson.ObjectIdHex(id), "meetings.meetingDate": bson.M{"$gte": startDate, "$lte": endDate}}},
		{"$group": bson.M{"_id": "$_id",
		 "meetings": bson.M{"$addToSet": "$meetings"}},
		},
	}
	pipe := c.Pipe(pipeline)
	result := bson.M{}
	err := pipe.One(&result)
	if err != nil {		
		return bson.M{}, err
	}
	return result, nil
}