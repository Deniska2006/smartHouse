package domain

import (
	"time"

	"github.com/google/uuid"
)

const SENSOR string = "SENSOR"
const ACTUATOR string = "ACTUATOR"

type Device struct {
	Id               uint64
	HouseId          uint64
	RoomId           uint64
	UUID             uuid.UUID
	SerialNumber     string
	Characteristics  *string
	Category         string
	Units            *string // for category SENSOR
	PowerConsumption *string // for category ACTUATOR
	CreatedDate      time.Time
	UpdatedDate      time.Time
	DeletedDate      *time.Time
}
