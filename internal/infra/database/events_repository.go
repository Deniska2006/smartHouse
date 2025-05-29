package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const EventsTableName = "events"

type event struct {
	Id          uint64    `db:"id,omitempty"`
	DeviceId    uint64    `db:"device_id"`
	RoomId      uint64    `db:"room_id"`
	Action      string    `db:"action"`
	CreatedDate time.Time `db:"created_date"`
}

type EventRepository interface {
	Save(e domain.Event) error
	Find(id uint64) (domain.Event, error)
}

type eventRepository struct {
	coll db.Collection
	sess db.Session
}

func NewEventRepository(sess db.Session) EventRepository {
	return eventRepository{
		coll: sess.Collection(EventsTableName),
		sess: sess,
	}
}

func (r eventRepository) Save(e domain.Event) error {
	event := r.mapDomainToModel(e)
	event.CreatedDate = time.Now()

	err := r.coll.InsertReturning(&event)
	if err != nil {
		return err
	}

	return nil
}

func (r eventRepository) Find(id uint64) (domain.Event, error) {
	var e event
	err := r.coll.
		Find(db.Cond{"id": id}).One(&e)
	if err != nil {
		return domain.Event{}, err
	}

	return r.mapModelToDomain(e), nil
}

func (r eventRepository) mapDomainToModel(e domain.Event) event {
	return event{
		Id:          e.Id,
		DeviceId:    e.DeviceId,
		RoomId:      e.RoomId,
		Action:      e.Action,
		CreatedDate: e.CreatedDate,
	}
}

func (r eventRepository) mapModelToDomain(e event) domain.Event {
	return domain.Event{
		Id:          e.Id,
		DeviceId:    e.DeviceId,
		RoomId:      e.RoomId,
		Action:      e.Action,
		CreatedDate: e.CreatedDate,
	}
}
