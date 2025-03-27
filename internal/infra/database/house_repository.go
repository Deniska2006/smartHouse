package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const HousesTableName = "houses"

type house struct {
	Id          uint64     `db:"id"`
	UserId      uint64     `db:"user_id"`
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	City        string     `db:"city"`
	Address     string     `db:"city"`
	Lat         float64    `db:"lat"`
	Lon         float64    `db:"lon"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type HouseRepository interface {
}

type houseRepository struct {
	coll db.Collection
	sess db.Session
}

func NewHouseRepository(sess db.Session) HouseRepository {
	return houseRepository{
		coll: sess.Collection(HousesTableName),
		sess: sess,
	}
}

func (r houseRepository) Save(h domain.House) (domain.House, error) {

}
