package models

import "gopkg.in/mgo.v2/bson"

// For new meeting creation
type NewMeeting struct {
	Id  bson.ObjectId `bson:"_id,omitempty" json:"id"`
	MeetingSubject  string `bson:"meetingSubject" json:"meetingSubject"`
	Meetingroom  string	`bson:"meetingroom" json:"meetingroom"`
	MeetingStartTime  string `bson:"meetingStartTime" json:"meetingStartTime"`
	MeetingEndTime  string `bson:"meetingEndTime" json:"meetingEndTime"`
	MeetingDate string `bson:"meetingDate" json:"meetingDate"`
	BookedBy string `bson:"bookedBy" json:"bookedBy"`
	MeetingStatus string `bson:"meetingStatus" json:"meetingStatus"`
	MeetingReccurance string `bson:"meetingReccurance" json:"meetingReccurance"`
}
