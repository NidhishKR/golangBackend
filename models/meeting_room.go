package models

import "gopkg.in/mgo.v2/bson"

type MeetingRoom struct {
	Id  bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name  string `bson:"name" json:"name"`
	Region  string`bson:"region" json:"region"`
	Building  string`bson:"building" json:"building"`
	CreatedBy  string `bson:"createdBy" json:"createdBy"`
	CreatedAt  string `bson:"createdAt" json:"createdAt"`
	Bookings []NewMeeting `bson:"bookings"`
}
