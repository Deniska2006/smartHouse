package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
)

type MeasurementRequest struct {
	// Id         uint64    `json:"id" validate:"required"`
	RoomId     uint64    `json:"roomId" validate:"required"`
	DeviceUUID uuid.UUID `json:"deviceUUID" validate:"required"`
	Value      string    `json:"value" validate:"required"`
}

func (r MeasurementRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		// Id:         r.Id,
		RoomId:     r.RoomId,
		DeviceUUID: r.DeviceUUID,
		Value:      r.Value,
	}, nil
}

