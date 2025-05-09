package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceRequest struct {
	SerialNumber     string  `json:"serialNumber" validate:"required"`
	Characteristics  *string `json:"characteristics"`
	Category         string  `json:"category" validate:"required"`
	Units            *string `json:"units"`
	PowerConsumption *string `json:"powerConsumption"`
}

type UpdateDeviceRequest struct {
	SerialNumber     *string  `json:"serialNumber"`
	Characteristics  *string `json:"characteristics"`
	Category         *string  `json:"category"`
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

func (r UpdateDeviceRequest) ToDomainModel() (interface{}, error) {
	var serialNumber,category string

	if r.SerialNumber != nil{
		serialNumber = *r.SerialNumber
	}
	if r.Category != nil{
		category = *r.Category
	}

	return domain.Device{
		SerialNumber:     serialNumber,
		Characteristics:  r.Characteristics,
		Category:         category,
		Units:            r.Units,
		PowerConsumption: r.PowerConsumption,
	}, nil
}
