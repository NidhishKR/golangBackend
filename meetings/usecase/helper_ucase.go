package usecase

import models "CleanArchMeetingRoom/models"

func checkTiming(a *meetingUsecase, m *models.NewMeeting, update bool) (bool,  error) {
	meetings, err := a.meetingRepos.GetConcurrentMeetings(m)
	for _, element := range meetings {
		if(update == true && m.Id == element.Id){
			continue
		}	else{
			if( element.MeetingStartTime == m.MeetingStartTime || element.MeetingEndTime == m.MeetingEndTime ){
				return true , err
			}
				if( m.MeetingStartTime > element.MeetingStartTime  && m.MeetingStartTime < element.MeetingEndTime){
				return true, err
			}
				if(m.MeetingEndTime > element.MeetingStartTime  && m.MeetingEndTime < element.MeetingEndTime){			
				return true, err
			}
			if(m.MeetingStartTime < element.MeetingStartTime  && m.MeetingEndTime > element.MeetingEndTime){			
				return true, err
			}
		}	
	}
	return false , err
}