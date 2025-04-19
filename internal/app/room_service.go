package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type RoomService interface {
	Save(rm domain.Room) (domain.Room, error)
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
