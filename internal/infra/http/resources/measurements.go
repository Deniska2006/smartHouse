package resources

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type MeasurementDto struct {
	Id       uint64 `json:"id"`
	RoomId   uint64 `json:"roomId"`
	DeviceId uint64 `json:"deviceId"`
	Value    string `json:"value"`
}

func (d MeasurementDto) DomainToDto(m domain.Measurement) MeasurementDto {
	return MeasurementDto{
		Id:       m.Id,
		RoomId:   m.RoomId,
		DeviceId: m.DeviceId,
		Value:    m.Value,
	}
}
