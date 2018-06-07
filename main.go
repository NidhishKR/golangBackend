package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"

	meetingsRepo "CleanArchMeetingRoom/meetings/repository"
	meetingsUcase "CleanArchMeetingRoom/meetings/usecase"
	httpDeliver "CleanArchMeetingRoom/meetings/delivery/http"
	userRepo "CleanArchMeetingRoom/user/repository"
	userUcase "CleanArchMeetingRoom/user/usecase"

	utils "CleanArchMeetingRoom/utils"
	middleware "CleanArchMeetingRoom/middleware"

)

func init() {
	filename := utils.GetFileName()
	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	dbHost := viper.GetString(`DATABASE.HOST`)
	dbName := viper.GetString(`DATABASE.NAME`)
	port := viper.GetString(`SERVER.PORT`)

	initMongo(dbHost)
	s := getSession()
	dbConn := s.DB(dbName)
	defer s.Close()
	app := gin.New()
	app.Use(gin.Logger())
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	ur := userRepo.NewMgoUserRepository(dbConn)
	uu := userUcase.NewUserUsecase(ur, timeoutContext)
	middL := middleware.InitMiddleware(uu)
	app.Use(middL.CORS())
	
	app.Use(middL.ValidateUser());

	mr := meetingsRepo.NewMgoMeetingsRepository(dbConn)
	mu := meetingsUcase.NewMeetingUsecase(mr, timeoutContext)
	httpDeliver.NewMeetingsHttpHandler(app, mu)

	app.Run(port)
}
