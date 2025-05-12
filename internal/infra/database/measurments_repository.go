package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

const MeasurementTableName = "measurements"

type measurement struct {
	Id          uint64    `db:"id,omitempty"`
	RoomId      uint64    `db:"room_id"`
	DeviceUUID  uuid.UUID `db:"device_uuid"`
	Value       string    `db:"value"`
	CreatedDate time.Time `db:"created_date"`
}

type MeasurementRepository interface {
	Save(m domain.Measurement) error

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



func (r measurementRepository) mapModelToDomain(m measurement) domain.Measurement {
	return domain.Measurement{
		Id:          m.Id,
		RoomId:      m.Id,
		DeviceUUID:  m.DeviceUUID,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
	}
}

func (r measurementRepository) mapDomainToModel(m domain.Measurement) measurement {
	return measurement{
		Id:          m.Id,
		RoomId:      m.RoomId,
		DeviceUUID:  m.DeviceUUID,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
	}
}
