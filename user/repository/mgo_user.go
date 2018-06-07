package repository

import (
	"context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	user "CleanArchMeetingRoom/user"
	models "CleanArchMeetingRoom/models"
	config "CleanArchMeetingRoom/config"
)

const USER string = "USER"

type mgoUserRepository struct {
	Conn *mgo.Database
}

func NewMgoUserRepository(Conn *mgo.Database) user.UserRepository {
	return &mgoUserRepository{Conn}
}

func (m *mgoUserRepository) GetUserByAuthToken(ctx context.Context, token string) (bool, error) {
	cn := config.CollectionNames(USER)
	c := m.Conn.C(cn)
	var user []models.User
	iter := c.Find(bson.M{"accessTokens.token": token}).Limit(500).Iter()
	err := iter.All(&user)
	if err != nil {
		return  false, err
	} else if len(user) == 0 {
		return  false, models.NOT_FOUND_ERROR
	}
	return true, nil
}