package http

import (
	"github.com/gin-gonic/gin"
	
	meetingsUcase "CleanArchMeetingRoom/meetings"
)

func NewMeetingsHttpHandler(app *gin.Engine, mu meetingsUcase.MeetingsUsecase) {
	handler := &HttpMeetingsHandler{
		MUsecase: mu,
	}
	app.GET("/meetingroom/region/:region", handler.GetByRegion)
	app.POST("/meetingroom/add", handler.AddMeetingroom)
}