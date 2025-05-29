package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const MeasurementTableName = "measurements"

type measurement struct {
	Id          uint64    `db:"id,omitempty"`
	DeviceId    uint64    `db:"device_id"`
	RoomId      uint64    `db:"room_id"`
	Value       string    `db:"value"`
	CreatedDate time.Time `db:"created_date"`
}

type MeasurementRepository interface {
	Save(m domain.Measurement) error
	Find(id uint64) (domain.Measurement, error)
}

type measurementRepository struct {
	coll db.Collection
	sess db.Session
}

func NewMeasurementRepository(sess db.Session) MeasurementRepository {
	return measurementRepository{
		coll: sess.Collection(MeasurementTableName),
		sess: sess,
	}
}

func (r measurementRepository) Save(m domain.Measurement) error {
	measurement := r.mapDomainToModel(m)
	measurement.CreatedDate = time.Now()

	err := r.coll.InsertReturning(&measurement)
	if err != nil {
		return err
	}

	return nil
}

func (r measurementRepository) Find(id uint64) (domain.Measurement, error) {
	var m measurement
	err := r.coll.
		Find(db.Cond{"id": id}).One(&m)
	if err != nil {
		return domain.Measurement{}, err
	}

	return r.mapModelToDomain(m), nil
}

func (r measurementRepository) mapDomainToModel(m domain.Measurement) measurement {
	return measurement{
		Id:          m.Id,
		DeviceId:    m.DeviceId,
		RoomId:      m.RoomId,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
	}
}

func (r measurementRepository) mapModelToDomain(m measurement) domain.Measurement {
	return domain.Measurement{
		Id:          m.Id,
		RoomId:      m.RoomId,
		DeviceId:    m.DeviceId,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
	}
}
