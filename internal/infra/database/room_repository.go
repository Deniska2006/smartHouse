package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomTableName = "rooms"

type room struct {
	Id          uint64     `db:"id,omitempty"`
	HouseId     uint64     `db:"house_id"`
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type RoomRepository interface {
	Save(rm domain.Room) (domain.Room, error)
	FindList(hId uint64) ([]domain.Room, error)
	Find(Id uint64) (domain.Room, error)
	Update(updt domain.Room, rm domain.Room)(domain.Room,error)
	Delete(rId uint64) error
}

type roomRepository struct {
	coll db.Collection
	sess db.Session
}

func NewRoomRepository(sess db.Session) RoomRepository {
	return roomRepository{
		coll: sess.Collection(RoomTableName),
		sess: sess,
	}
}

func (r roomRepository) Save(rm domain.Room) (domain.Room, error) {
	room := r.mapDomainToModel(rm)
	room.CreatedDate = time.Now()
	room.UpdatedDate = time.Now()

	err := r.coll.InsertReturning(&room)
	if err != nil {
		return domain.Room{}, err
	}

	return r.mapModelToDomain(room), nil
}

func (r roomRepository) FindList(hId uint64) ([]domain.Room, error) {
	var rooms []room
	err := r.coll.
		Find(db.Cond{
			"house_id":     hId,
			"deleted_date": nil}).All(&rooms)
	if err != nil {
		return nil, err
	}

	return r.mapModelToDomainCollection(rooms), nil
}

func (r roomRepository) Find(Id uint64) (domain.Room, error) {
	var rm room
	err := r.coll.
		Find(db.Cond{
			"id":           Id,
			"deleted_date": nil}).One(&rm)
	if err != nil {
		return domain.Room{}, err
	}

	return r.mapModelToDomain(rm), nil
}

func (r roomRepository) Update(updt domain.Room, rm domain.Room)(domain.Room,error) {
	err := r.coll.Find(db.Cond{"id": rm.Id}).Update(r.mapDomainToModelUpdate(updt))
	if err != nil {
		return domain.Room{}, err
	}

	return r.mapModelToDomainUpdate(updt,rm), nil
}

func (r roomRepository) Delete(rId uint64) error {
	err := r.coll.Find(db.Cond{"id": rId}).Update(map[string]interface{}{
		"deleted_date" : time.Now(),
	})
	if err != nil {
		return  err
	}

	return nil
}

func (r roomRepository) mapModelToDomainUpdate(updt domain.Room, rm domain.Room) domain.Room {
	if updt.Name != rm.Name && updt.Name != "" {
		rm.Name = updt.Name
	}
	if updt.Description != rm.Description && updt.Description != nil {
		rm.Description = updt.Description
	}
	rm.UpdatedDate = time.Now()
	return rm
}

func (r roomRepository) mapDomainToModelUpdate(updt domain.Room) map[string]interface{} {
	result := make(map[string]interface{},1)
	if updt.Name != "" {
		result["name"] = updt.Name
	}
	if updt.Description != nil {
		result["description"] = updt.Description
	}
	result["updated_date"] = time.Now()
	return result
}

func (r roomRepository) mapModelToDomainCollection(rooms []room) []domain.Room {
	rs := make([]domain.Room, len(rooms))
	for i, rm := range rooms {
		rs[i] = r.mapModelToDomain(rm)
	}
	return rs
}

func (r roomRepository) mapDomainToModel(rm domain.Room) room {
	return room{
		Id:          rm.Id,
		HouseId:     rm.HouseId,
		Name:        rm.Name,
		Description: rm.Description,
		CreatedDate: rm.CreatedDate,
		UpdatedDate: rm.UpdatedDate,
	}
}

func (r roomRepository) mapModelToDomain(room room) domain.Room {
	return domain.Room{
		Id:          room.Id,
		HouseId:     room.HouseId,
		Name:        room.Name,
		Description: room.Description,
		CreatedDate: room.CreatedDate,
		UpdatedDate: room.UpdatedDate,
	}
}
