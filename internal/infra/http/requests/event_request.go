package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
)

type EventRequest struct {
	DeviceUUID uuid.UUID `json:"deviceUUID" validate:"required"`
	Action     string    `json:"action" validate:"required"`
}

func (r EventRequest) ToDomainModel() (interface{}, error) {
	return domain.Event{
		DeviceUUID: r.DeviceUUID,
		Action:     r.Action,
	}, nil
}
