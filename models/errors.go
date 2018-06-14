package models

import "errors"

var (
	INTERNAL_SERVER_ERROR =  errors.New("Internal Server Error")
	NOT_FOUND_ERROR       =  errors.New("Your requested Item is not found")
	CONFLIT_ERROR         =  errors.New("Your Item already exist")
	TIME_SLOT_BOOKED		  =  errors.New("Meeting room already booked")
	NO_USER_FOUND		  =  errors.New("No user found")
	INVALID_TIME_INPUT 		  =  errors.New("lnvalid meeting time")
	INVALID_INPUT         =  errors.New("lnvalid input")
)
