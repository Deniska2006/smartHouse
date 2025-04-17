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

type HouseController struct {
	houseService app.HouseService
}

func NewHouseController(hs app.HouseService) HouseController {
	return HouseController{
		houseService: hs,
	}
}

func (c HouseController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		house, err := requests.Bind(r, requests.HouseRequest{}, domain.House{})
		if err != nil {
			log.Printf("HouseController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		house.UserId = user.Id

		house, err = c.houseService.Save(house)
		if err != nil {
			log.Printf("HouseController.Save(c.houseService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		var houseDto resources.HouseDto
		houseDto = houseDto.DomainToDto(house)
		Success(w, houseDto)
	}
}

func (c HouseController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updt, err := requests.Bind(r, requests.UpdateHouseRequest{}, domain.House{})
		if err != nil {
			log.Printf("HouseController.Update(requests.BindToMap): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		house := r.Context().Value(HouseKey).(domain.House)

		house, err = c.houseService.Update(updt, house)
		if err != nil {
			log.Printf("HouseController.Update(c.houseService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		var houseDto resources.HouseDto
		houseDto = houseDto.DomainToDto(house)
		Success(w, houseDto)
	}
}

func (c HouseController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		house := r.Context().Value(HouseKey).(domain.House)

		err := c.houseService.Delete(house.Id)
		if err != nil {
			log.Printf("HouseController.Delete(c.houseService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.Message{Response: "House was deleted"})
	}
}

func (c HouseController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		house := r.Context().Value(HouseKey).(domain.House)
		user := r.Context().Value(UserKey).(domain.User)

		if house.UserId != user.Id {
			err := errors.New("Acces denied")
			Forbidden(w, err)
			return
		}

		var houseDto resources.HouseDto
		houseDto = houseDto.DomainToDto(house)
		Success(w, houseDto)
	}
}

func (c HouseController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		houses, err := c.houseService.FindList(user.Id)
		if err != nil {
			log.Printf("(HouseController) FindList( c.houseService.FindList(user.Id)): %s", err)
			InternalServerError(w, err)
		}

		Success(w, resources.HouseDto{}.DomainToDtoCollection(houses))
	}
}
