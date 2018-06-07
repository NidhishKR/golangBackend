package meetings

import (
	"CleanArchMeetingRoom/models"
	
	"context"
)

type MeetingsRepository interface {
	GetByRegion(ctx context.Context, id string) (*[]models.MeetingRoom, error)
}
