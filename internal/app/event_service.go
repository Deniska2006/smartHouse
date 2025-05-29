package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/google/uuid"
)

type EventService interface {
	Save(e domain.Event) error
	FindDeviceByUUID(u uuid.UUID) (domain.Device, error)
	Find(id uint64) (interface{}, error)
}

type eventService struct {
	eventRepo  database.EventRepository
	deviceRepo database.DeviceRepository
}

func NewEventtService(er database.EventRepository, dr database.DeviceRepository) EventService {
	return eventService{
		eventRepo:  er,
		deviceRepo: dr,
	}
}

func (s eventService) Save(e domain.Event) error {
	err := s.eventRepo.Save(e)
	if err != nil {
		log.Printf("eventService.Save(s.eventRepo.Save): %s", err)
		return err
	}

	return nil
}

func (s eventService) FindDeviceByUUID(u uuid.UUID) (domain.Device, error) {
	device, err := s.deviceRepo.FindDeviceByUUID(u)
	if err != nil {
		log.Printf("eventService.FindDeviceByUUID(s.deviceRepo.FindDeviceByUUID): %s", err)
		return domain.Device{}, err
	}

	return device, nil
}

func (s eventService) Find(id uint64) (interface{}, error) {
	event, err := s.eventRepo.Find(id)
	if err != nil {
		log.Printf("eventService.FindDeviceByUUID(s.eventRepo.Find): %s", err)
		return domain.Event{}, err
	}

	return event, nil
}
