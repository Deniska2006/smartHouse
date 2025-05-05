package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
)

type DeviceDto struct {
	Id               uint64     `json:"id"`
	HouseId          uint64     `json:"houseId"`
	RoomId           uint64     `json:"roomId"`
	UUID             uuid.UUID  `json:"uuid"`
	SerialNumber     string     `json:"serialNumber"`
	Characteristics  *string    `json:"characteristics,omitempty"`
	Category         string     `json:"category"`
	Units            *string    `json:"units,omitempty"`
	PowerConsumption *string    `json:"power_consumption,omitempty"`
	CreatedDate      time.Time  `json:"createdDate"`
	UpdatedDate      time.Time  `json:"updatedDate"`
	DeletedDate      *time.Time `json:"deletedDate,omitempty"`
}

func (d DeviceDto) DomainToDto(r domain.Device) DeviceDto {
	return DeviceDto{
		Id:               r.Id,
		HouseId:          r.HouseId,
		RoomId:           r.RoomId,
		UUID:             r.UUID,
		SerialNumber:     r.SerialNumber,
		Characteristics:  r.Characteristics,
		Category:         r.Category,
		Units:            r.Units,
		PowerConsumption: r.PowerConsumption,
		CreatedDate:      r.CreatedDate,
		UpdatedDate:      r.UpdatedDate,
		DeletedDate:      r.DeletedDate,
	}
}
