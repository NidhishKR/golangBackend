package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserId  string `bson:"userId" json:"userId"`
	EmailId string `bson:"emailId" json:"emailId"`
	FirstName string`bson:"firstName" json:"firstName"`
	LastName string`bson:"lastName" json:"lastName"`
}