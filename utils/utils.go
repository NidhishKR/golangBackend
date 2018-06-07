package utils

import (
	"os"
	"strings"
	validator "gopkg.in/go-playground/validator.v9"
	models "CleanArchMeetingRoom/models"
)

func GetFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		panic("Please specify Environment")
	} else {
		filename := []string{"config.", env, ".json"}
		FilePath := strings.Join(filename, "")
		return FilePath
	}
}

func IsValidMR(m *models.MeetingRoom) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func IsValidM(m *models.NewMeeting) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}