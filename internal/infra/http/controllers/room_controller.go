package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type RoomController struct {
	roomService app.RoomService
}

func NewRoomController(rs app.RoomService) RoomController {
	return RoomController{
		roomService: rs,
	}
}

func (c RoomController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		room, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		house := r.Context().Value(HouseKey).(domain.House)
		room.HouseId = house.Id

		room, err = c.roomService.Save(room)
		if err != nil {
			log.Printf("RoomController.Save(c.roomService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		var roomDto resources.RoomDto
		roomDto = roomDto.DomainToDto(room)
		Success(w, roomDto)
	}
}
