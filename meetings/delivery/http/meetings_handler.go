package http

import (
	"github.com/gin-gonic/gin"
	"net/http"

	meetingsUcase "CleanArchMeetingRoom/meetings"
	models "CleanArchMeetingRoom/models"
	utils "CleanArchMeetingRoom/utils"
)

type HttpMeetingsHandler struct {
	MUsecase meetingsUcase.MeetingsUsecase
}

type ResponseError struct {
	Message string `json:"message"`
}

func (a *HttpMeetingsHandler) GetByRegion(c *gin.Context) {
	region := c.Param("region")
	m, err := a.MUsecase.GetByRegion(c, region)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.JSON(http.StatusOK, m)
}

func (a *HttpMeetingsHandler) AddMeetingroom(c *gin.Context) {
	var meetingRoom models.MeetingRoom

	err := c.Bind(&meetingRoom)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := utils.IsValidMR(&meetingRoom); !ok {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	mr, err := a.MUsecase.AddMeetingroom(c, &meetingRoom)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	
	c.JSON(http.StatusCreated, mr)
}

func (a *HttpMeetingsHandler) AddMeeting(c *gin.Context) {
	var meeting models.NewMeeting

	err := c.Bind(&meeting)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := utils.IsValidM(&meeting); !ok {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	if !utils.IsValidString(meeting.BookedBy) {
		c.JSON(http.StatusNoContent, models.INVALID_INPUT)
	}
	
	m, err := a.MUsecase.AddMeeting(c, &meeting)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	
	c.JSON(http.StatusCreated, m)
}


func (a *HttpMeetingsHandler) GetMeetingsByDateRange(c *gin.Context) {
	startDate := c.Param("startDate")
	endDate := c.Param("endDate")
	meetingroomId := c.Param("meetingroom")
    if (!utils.IsValidString(startDate) || !utils.IsValidString(endDate) || !utils.IsValidString(meetingroomId)) {
		c.JSON(http.StatusNoContent, models.INVALID_INPUT)
	}
	m, err := a.MUsecase.GetMeetingsByDateRange(c, startDate, endDate, meetingroomId)
	if err != nil && err.Error() == "not found" {
		p := []models.NewMeeting{}
		c.JSON(http.StatusOK, gin.H{
				"status":  err.Error(),
				"code":    "400",
				"meetings": p,
			})
	}
	c.JSON(http.StatusCreated, m)
}
