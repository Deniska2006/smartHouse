package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type DeviceService interface {
	Save(h domain.Device) (domain.Device, error)
	FindList(rId uint64) ([]domain.Device, error)
	Find(id uint64) (interface{}, error)
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

func (s deviceService) FindList(rId uint64) ([]domain.Device, error) {
	devices, err := s.deviceRepo.FindList(rId)
	if err != nil {
		log.Printf("deviceService.FindList(s.deviceRepo.FindList): %s", err)
		return nil, err
	}

	return devices, nil
}

func (s deviceService) Find(id uint64) (interface{}, error) {

	device, err := s.deviceRepo.Find(id)
	if err != nil {
		log.Printf("houseService.Find(s.houseRepo.Find): %s", err)
		return domain.Device{}, err
	}

	return device, nil
}