package repository

import (
	"context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	user "CleanArchMeetingRoom/user"
	models "CleanArchMeetingRoom/models"
	"fmt"
)

type mgoUserRepository struct {
	Conn *mgo.Database
}

func NewMgoUserRepository(Conn *mgo.Database) user.UserRepository {
	return &mgoUserRepository{Conn}
}

func (m *mgoUserRepository) GetUserByAuthToken(ctx context.Context, token string) (bool, error) {
	cn := models.USERS.DB.MODELS.COLLECTION
	c := m.Conn.C(cn)
	var user []models.User
	iter := c.Find(bson.M{"accessTokens.token": token}).Limit(500).Iter()
	err := iter.All(&user)
	if err != nil {
		return  false, models.NOT_FOUND_ERROR
	} else if len(user) == 0 {
		return  false, models.NOT_FOUND_ERROR
	}
	return true, nil
}

func (m *mgoUserRepository) GetUserByEmailId(ctx context.Context, emailId string) (models.GlobalUser, error) {
	cn := models.GLOBALUSER.DB.MODELS.COLLECTION
	c := m.Conn.C(cn)
	var user models.GlobalUser
	err := c.Find(bson.M{"email": emailId}).One(&user)
	if err != nil {
		return models.GlobalUser{}, err
	}
	return user, nil
}