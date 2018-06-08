package models

type ModelsDetails struct {
    NAME string
    COLLECTION string
    STRUCT string
}

type Models struct {
	MODELS *ModelsDetails
}
type constent struct {
	DB *Models 
}
var (
	MEETINGROOM = constent{DB: &Models{MODELS: &ModelsDetails{NAME: "MeetingRoom", COLLECTION: "meetingroom", STRUCT: "MeetingRoom"}}}
	MEETING = constent{DB: &Models{MODELS: &ModelsDetails{NAME: "Meeting", COLLECTION: "meetingroom", STRUCT: "NewMeeting"}}}
	USERS = constent{DB: &Models{MODELS: &ModelsDetails{NAME: "User", COLLECTION: "users", STRUCT: "User"}}}
)