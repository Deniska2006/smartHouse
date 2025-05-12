package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/google/uuid"
)

type MeasurementService interface {
	Save(m domain.Measurement) error
	FindDeviceByUUID(u uuid.UUID) error
}

type measurementService struct {
	measRepo   database.MeasurementRepository
	deviceRepo database.DeviceRepository
}

func NewMeasurementService(mr database.MeasurementRepository, dr database.DeviceRepository) MeasurementService {
	return measurementService{
		measRepo:   mr,
		deviceRepo: dr,
	}
}

func (s measurementService) Save(m domain.Measurement) error {
	err := s.measRepo.Save(m)
	if err != nil {
		log.Printf("measurementService.Save(s.measRepo.Save): %s", err)
		return err
	}

	return nil
}

func (s measurementService) FindDeviceByUUID(u uuid.UUID) error {
	err := s.deviceRepo.FindDeviceByUUID(u)
	if err != nil {
		log.Printf("measurementService.FindDeviceByUUID(s.measRepo.FindDeviceByUUID): %s", err)
		return err
	}

	return nil
}
