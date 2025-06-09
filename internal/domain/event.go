package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id          uint64
	RoomId      uint64
	DeviceId    uint64
	DeviceUUID  uuid.UUID	
	Action      string
	CreatedDate time.Time
}
