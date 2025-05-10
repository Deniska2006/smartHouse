package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type RoomService interface {
	Save(rm domain.Room) (domain.Room, error)
	FindList(hId uint64) ([]domain.Room, error)
	Find(id uint64) (interface{}, error)
	Update(updt domain.Room, rm domain.Room) (domain.Room, error)
	Delete(rId uint64) error
	FindDevices(rid uint64) ([]domain.Device, error)
}

type roomService struct {
	roomRepo   database.RoomRepository
	deviceRepo database.DeviceRepository
}

func NewRoomService(rr database.RoomRepository, dr database.DeviceRepository) RoomService {
	return roomService{
		roomRepo:   rr,
		deviceRepo: dr,
	}
}

func (s roomService) Save(rm domain.Room) (domain.Room, error) {
	room, err := s.roomRepo.Save(rm)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Room{}, err
	}

	return room, nil
}

func (s roomService) FindList(hId uint64) ([]domain.Room, error) {
	rooms, err := s.roomRepo.FindList(hId)
	if err != nil {
		log.Printf("roomService.FindList(s.roomRepo.FindList): %s", err)
		return nil, err
	}

	return rooms, nil
}

func (s roomService) Find(id uint64) (interface{}, error) {
	room, err := s.roomRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.roomRepo.Find): %s", err)
		return domain.House{}, err
	}

	return room, nil
}

func (s roomService) FindDevices(rid uint64) ([]domain.Device, error) {
	devices, err := s.deviceRepo.FindList(rid)
	if err != nil {
		log.Printf("roomService.FindDevices(s.deviceRepo.FindList): %s", err)
		return nil, err
	}

	return devices, nil
}

func (s roomService) Update(updt domain.Room, rm domain.Room) (domain.Room, error) {
	room, err := s.roomRepo.Update(updt, rm)
	if err != nil {
		log.Printf("roomService.Update(s.roomRepo.Update): %s", err)
		return domain.Room{}, err
	}

	return room, nil
}

func (s roomService) Delete(rId uint64) error {
	err := s.roomRepo.Delete(rId)
	if err != nil {
		log.Printf("roomService.Delete(s.roomRepo.Delete): %s", err)
		return err
	}

	return nil
}
