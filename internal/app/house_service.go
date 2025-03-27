package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type HouseService interface {
	Save(h domain.House) (domain.House, error)
}

type houseService struct {
	houseRepo database.HouseRepository
}

func NewHouseService(hr database.HouseRepository) HouseService {
	return houseService{
		houseRepo: hr,
	}
}

func (s houseService) Save(h domain.House) (domain.House, error) {
	house, err := s.houseRepo.Save(h)
	if err != nil {
		log.Printf("houseService.Save(s.houseRepo.Save): %s", err)
		return domain.House{}, err
	}

	return house, nil
}
