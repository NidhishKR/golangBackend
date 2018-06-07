package http

import (
	"github.com/gin-gonic/gin"
	"net/http"

	meetingsUcase "CleanArchMeetingRoom/meetings"
	models "CleanArchMeetingRoom/models"
	validator "gopkg.in/go-playground/validator.v9"
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

	if ok, err := isRequestValid(&meetingRoom); !ok {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	mr, err := a.MUsecase.AddMeetingroom(c, &meetingRoom)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	
	c.JSON(http.StatusCreated, mr)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func isRequestValid(m *models.MeetingRoom) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}