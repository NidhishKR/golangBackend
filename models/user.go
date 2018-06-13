package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id  bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserId  string `bson:"userId" json:"userId"`
	EmailId string `bson:"emailId" json:"emailId"`
	FirstName string`bson:"firstName" json:"firstName"`
	LastName string`bson:"lastName" json:"lastName"`
}

type GlobalUser struct {
	EmployeeName string`bson:"employee_name" json:"employee_name"`
	UserId  string `bson:"employee_id" json:"employee_id"`
	EmailId string `bson:"email" json:"email"`
	ProfileImageUrl string`bson:"profile_img" json:"profile_img"`
}