package main

import (
	"log"
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

func initMongo(host string) bool {
	url := host
	log.Println("Establishing MongoDB connection...")
	var err error
	Session, err = mgo.Dial(url)
	if err != nil {
		log.Fatal("Cannot connect to MongoDB!")
		return true
	} else {
		log.Println("Connected to ", url)
		return false
	}
}

func getSession() mgo.Session {
	return *Session.Copy()
}
