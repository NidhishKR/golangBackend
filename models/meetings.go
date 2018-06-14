package models

import "gopkg.in/mgo.v2/bson"
// For new meeting creation
type NewMeeting struct {
	Id  string 	`bson:"id" json:"meetingId"`
	MeetingSubject  string `bson:"meetingSubject" json:"meetingSubject"`
	Meetingroom  string	`bson:"meetingroom" json:"meetingroom"`
	MeetingStartTime  string `bson:"meetingStartTime" json:"meetingStartTime"`
	MeetingEndTime  string `bson:"meetingEndTime" json:"meetingEndTime"`
	MeetingDate string `bson:"meetingDate" json:"meetingDate"`
	BookedBy string `bson:"bookedBy" json:"bookedBy"`
	UserDetails  *GlobalUser `bson:"userDetails" json:"userDetails"`
	MeetingStatus string `bson:"meetingStatus" json:"meetingStatus"`
	MeetingReccurance int `bson:"meetingReccurance" json:"meetingReccurance"`
	ReccuranceCount ReccuranceCount `bson:"reccuranceCount" json:"reccuranceCount"`
}

type ReccuranceCount struct {
	Id  string `bson:"id" json:"id"`
	Repetition  int `json:"repetition"`
}

type MeetingRes struct{
	Id  bson.ObjectId   `json:"_id" bson:"_id"`
	Meetings  [] NewMeeting    `json:"meetings" bson:"meetings"`
}
