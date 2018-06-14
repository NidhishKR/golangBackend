package http

import (
	"net/http"

	models "CleanArchMeetingRoom/models"
)

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