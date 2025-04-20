package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type RoomDto struct {
	Id          uint64     `json:"id"`
	HouseId     uint64     `json:"house_Id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeletedDate *time.Time `json:"deletedDate,omitempty"`
}

func (d RoomDto) DomainToDto(r domain.Room) RoomDto {
	return RoomDto{
		Id:          r.Id,
		HouseId:     r.HouseId,
		Name:        r.Name,
		Description: r.Description,
		CreatedDate: r.CreatedDate,
		UpdatedDate: r.UpdatedDate,
		DeletedDate: r.DeletedDate,
	}
}

func (d RoomDto) DomainToDtoCollection(rooms []domain.Room) []RoomDto {
	rs := make([]RoomDto, len(rooms))
	for i, room := range rooms {
		rs[i] = d.DomainToDto(room)
	}
	return rs
}
