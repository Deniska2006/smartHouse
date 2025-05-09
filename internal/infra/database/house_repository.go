package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const HousesTableName = "houses"

type house struct {
	Id          uint64     `db:"id,omitempty"`
	UserId      uint64     `db:"user_id"`
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	City        string     `db:"city"`
	Address     string     `db:"address"`
	Lat         float64    `db:"lat"`
	Lon         float64    `db:"lon"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type HouseRepository interface {
	Save(h domain.House) (domain.House, error)
	FindList(uId uint64) ([]domain.House, error)
	Find(id uint64) (domain.House, error)
	Update(updt domain.House, h domain.House) (domain.House, error)
	Delete(hId uint64) error
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
	hs := r.mapDomainToModel(h)
	hs.CreatedDate = time.Now()
	hs.UpdatedDate = time.Now()

	err := r.coll.InsertReturning(&hs)
	if err != nil {
		return domain.House{}, err
	}

	h = r.mapModelToDomain(hs)
	return h, nil
}

func (r houseRepository) Update(updt domain.House, h domain.House) (domain.House, error) {
	err := r.coll.Find(db.Cond{"id": h.Id}).Update(r.mapDomainToInterfaceUpdate(updt))
	if err != nil {
		return domain.House{}, err
	}

	return r.mapUpdateToDomain(updt, h), nil
}

func (r houseRepository) Delete(hId uint64) error {
	err := r.coll.Find(db.Cond{"id": hId}).Update(map[string]interface{}{
		"deleted_date": time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r houseRepository) FindList(uId uint64) ([]domain.House, error) {
	var houses []house
	err := r.coll.
		Find(db.Cond{
			"user_id":      uId,
			"deleted_date": nil}).All(&houses)
	if err != nil {
		return nil, err
	}

	hs := r.mapModelToDomainCollection(houses)
	return hs, nil
}

func (r houseRepository) Find(id uint64) (domain.House, error) {
	var h house
	err := r.coll.
		Find(db.Cond{
			"id":           id,
			"deleted_date": nil}).One(&h)
	if err != nil {
		return domain.House{}, err
	}

	hs := r.mapModelToDomain(h)
	return hs, nil
}

func (r houseRepository) mapDomainToModel(d domain.House) house {
	return house{
		Id:          d.Id,
		UserId:      d.UserId,
		Name:        d.Name,
		Description: d.Description,
		City:        d.City,
		Address:     d.Address,
		Lat:         d.Lat,
		Lon:         d.Lon,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r houseRepository) mapDomainToInterfaceUpdate(updt domain.House) map[string]interface{} {
	result := make(map[string]interface{}, 1)
	if updt.Name != "" {
		result["name"] = updt.Name
	}
	if updt.Description != nil {
		result["description"] = updt.Description
	}
	if updt.City != "" {
		result["city"] = updt.City
	}
	if updt.Address != "" {
		result["address"] = updt.Address
	}
	if updt.Lat != 0 {
		result["lat"] = updt.Lat
	}
	if updt.Lon != 0 {
		result["lon"] = updt.Lon
	}
	result["updated_date"] = time.Now()
	return result
}

func (r houseRepository) mapModelToDomain(d house) domain.House {
	return domain.House{
		Id:          d.Id,
		UserId:      d.UserId,
		Name:        d.Name,
		Description: d.Description,
		City:        d.City,
		Address:     d.Address,
		Lat:         d.Lat,
		Lon:         d.Lon,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r houseRepository) mapModelToDomainCollection(houses []house) []domain.House {
	hs := make([]domain.House, len(houses))
	for i, h := range houses {
		hs[i] = r.mapModelToDomain(h)
	}

	return hs
}

func (r houseRepository) mapUpdateToDomain(updt domain.House, h domain.House) domain.House {
	if updt.Name != h.Name && updt.Name != "" {
		h.Name = updt.Name
	}
	if updt.Description != h.Description && updt.Description != nil {
		h.Description = updt.Description
	}
	if updt.City != h.City && updt.City != "" {
		h.City = updt.City
	}
	if updt.Address != h.Address && updt.Address != "" {
		h.Address = updt.Address
	}
	if updt.Lat != h.Lat && updt.Lat != 0 {
		h.Lat = updt.Lat
	}
	if updt.Lon != h.Lon && updt.Lon != 0 {
		h.Lon = updt.Lon
	}
	h.UpdatedDate = time.Now()
	return h
}
