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
}

type roomService struct {
	roomRepo database.RoomRepository
}

func NewRoomService(rr database.RoomRepository) RoomService {
	return roomService{
		roomRepo: rr,
	}
}

func (s roomService) Save(rm domain.Room) (domain.Room, error) {
	house, err := s.roomRepo.Save(rm)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Room{}, err
	}

	return house, nil
}

func (s roomService) FindList(hId uint64) ([]domain.Room, error) {
	rooms, err := s.roomRepo.FindList(hId)
	if err != nil {
		log.Printf("roomService.FindList(s.roomRepo.FindList): %s", err)
		return nil, err
	}

	return rooms, nil
}

// func (s roomService) Find(Id uint64) (domain.Room, error) {
// 	room, err := s.roomRepo.Find(Id)
// 	if err != nil {
// 		log.Printf("roomService.Find(s.roomRepo.Find): %s", err)
// 		return domain.Room{}, err
// 	}

// 	return room, nil
// }

func (s roomService) Find(id uint64) (interface{}, error) {
	house, err := s.roomRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.roomRepo.Find): %s", err)
		return domain.House{}, err
	}

	return house, nil
}
