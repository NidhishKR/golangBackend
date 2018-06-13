package models

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
	MeetingReccurance string `bson:"meetingReccurance" json:"meetingReccurance"`
}
