package http

import (
	"github.com/gin-gonic/gin"
	"net/http"

	meetingsUcase "CleanArchMeetingRoom/meetings"
	models "CleanArchMeetingRoom/models"
)

type HttpMeetingsHandler struct {
	MUsecase meetingsUcase.MeetingsUsecase
}

type ResponseError struct {
	Message string `json:"message"`
}

func (a *HttpMeetingsHandler) GetByRegion(c *gin.Context) {
	region := c.Param("region")
	art, err := a.MUsecase.GetByRegion(c, region)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.JSON(http.StatusOK, art)
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