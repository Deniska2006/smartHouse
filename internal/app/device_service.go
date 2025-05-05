package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type DeviceService interface {
	Save(h domain.Device) (domain.Device, error)
}

type deviceService struct {
	deviceRepo database.DeviceRepository
}

func NewDeviceService(dr database.DeviceRepository) DeviceService {
	return deviceService{
		deviceRepo: dr,
	}
}

func (s deviceService) Save(h domain.Device) (domain.Device, error) {
	house, err := s.deviceRepo.Save(h)
	if err != nil {
		log.Printf("deviceService.Save(s.deviceRepo.Save): %s", err)
		return domain.Device{}, err
	}

	return house, nil
}
