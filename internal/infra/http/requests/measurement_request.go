package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
)

type MeasurementRequest struct {

	DeviceUUID uuid.UUID `json:"deviceUUID" validate:"required"`
	Value      string    `json:"value" validate:"required"`
}

func (r MeasurementRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		DeviceUUID: r.DeviceUUID,
		Value:      r.Value,
	}, nil
}

