package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type HouseDto struct {
	Id          uint64     `json:"id"`
	UserId      uint64     `json:"userId"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	City        string     `json:"city"`
	Address     string     `json:"address"`
	Lat         float64    `json:"lat"`
	Lon         float64    `json:"lon"`
	Rooms       []RoomDto  `json:"rooms"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeletedDate *time.Time `json:"deletedDate,omitempty"`
}

type Message struct {
	Response string
}

func (d HouseDto) DomainToDto(h domain.House) HouseDto {
	return HouseDto{
		Id:          h.Id,
		UserId:      h.UserId,
		Name:        h.Name,
		Description: h.Description,
		City:        h.City,
		Address:     h.Address,
		Lat:         h.Lat,
		Lon:         h.Lon,
		Rooms:       RoomDto{}.DomainToDtoCollection(h.Rooms),
		CreatedDate: h.CreatedDate,
		UpdatedDate: h.UpdatedDate,
		DeletedDate: h.DeletedDate,
	}
}

func (d HouseDto) DomainToDtoCollection(houses []domain.House) []HouseDto {
	hs := make([]HouseDto, len(houses))
	for i, house := range houses {
		hs[i] = d.DomainToDto(house)
	}
	return hs
}
