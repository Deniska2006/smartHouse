package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type EventDto struct {
	Id       uint64 `json:"id"`
	RoomId   uint64 `json:"roomId"`
	DeviceId uint64 `json:"deviceId"`
	Action   string `json:"action"`
}

func (d EventDto) DomainToDto(m domain.Event) EventDto {
	return EventDto{
		Id:       m.Id,
		RoomId:   m.RoomId,
		DeviceId: m.DeviceId,
		Action:   m.Action,
	}
}
