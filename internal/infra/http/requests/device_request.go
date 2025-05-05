package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceRequest struct {
	SerialNumber     string  `json:"serialNumber" validate:"required"`
	Characteristics  *string `json:"characteristics"`
	Category         string  `json:"category" validate:"required"`
	Units            *string `json:"units"`
	PowerConsumption *string `json:"powerConsumption"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		SerialNumber:     r.SerialNumber,
		Characteristics:  r.Characteristics,
		Category:         r.Category,
		Units:            r.Units,
		PowerConsumption: r.PowerConsumption,
	}, nil
}
