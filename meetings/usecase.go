package meetings

import (
	"context"
	
	model "CleanArchMeetingRoom/models"
)

type MeetingsUsecase interface {
	GetByRegion(ctx context.Context, region string) (*[]model.MeetingRoom, error)
}
