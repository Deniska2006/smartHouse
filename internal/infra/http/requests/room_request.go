package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type RoomRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type UpdateRoomRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (r RoomRequest) ToDomainModel() (interface{}, error) {
	return domain.Room{
		Name:        r.Name,
		Description: r.Description,
	}, nil
}

func (r UpdateRoomRequest) ToDomainModel() (interface{}, error) {
	var name string
	if r.Name != nil {
		name = *r.Name
	}

	return domain.Room{
		Name:        name,
		Description: r.Description,
	}, nil
}
