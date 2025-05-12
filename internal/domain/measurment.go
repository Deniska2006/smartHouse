package domain

import (
	"time"

	"github.com/google/uuid"
)

type Measurement struct {
	Id          uint64 // будемо автоінкрементуватись в скетчі
	RoomId      uint64
	DeviceUUID  uuid.UUID
	Value       string
	CreatedDate time.Time
}

// Чи треба нам залишати  DeviceId  RoomId бо ми будемо передавати uuid
