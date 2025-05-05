package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

const DeviceTableName = "devices"

type device struct {
	Id               uint64     `db:"id,omitempty"`
	HouseId          uint64     `db:"house_id"`
	RoomId           uint64     `db:"room_id"`
	UUID             uuid.UUID  `db:"uuid"`
	SerialNumber     string     `db:"serial_number"`
	Characteristics  *string    `db:"characteristics"`
	Category         string     `db:"category"`
	Units            *string    `db:"units"`
	PowerConsumption *string    `db:"power_consumption"`
	CreatedDate      time.Time  `db:"created_date"`
	UpdatedDate      time.Time  `db:"updated_date"`
	DeletedDate      *time.Time `db:"deleted_date"`
}

type DeviceRepository interface {
	Save(d domain.Device) (domain.Device, error)
}

type deviceRepository struct {
	coll db.Collection
	sess db.Session
}

func NewDeviceRepository(sess db.Session) DeviceRepository {
	return deviceRepository{
		coll: sess.Collection(DeviceTableName),
		sess: sess,
	}
}

func (r deviceRepository) Save(d domain.Device) (domain.Device, error) {
	dvc := r.mapDomainToModel(d)
	dvc.CreatedDate = time.Now()
	dvc.UpdatedDate = time.Now()

	err := r.coll.InsertReturning(&dvc)
	if err != nil {
		return domain.Device{}, err
	}

	return r.mapModelToDomain(dvc), nil
}

func (r deviceRepository) mapDomainToModel(d domain.Device) device {
	return device{
		Id:               d.Id,
		HouseId:          d.HouseId,
		RoomId:           d.RoomId,
		UUID:             d.UUID,
		SerialNumber:     d.SerialNumber,
		Characteristics:  d.Characteristics,
		Category:         d.Category,
		Units:            d.Units,
		PowerConsumption: d.PowerConsumption,
	}
}

func (r deviceRepository) mapModelToDomain(d device) domain.Device {
	return domain.Device{
		Id:               d.Id,
		HouseId:          d.HouseId,
		RoomId:           d.RoomId,
		UUID:             d.UUID,
		SerialNumber:     d.SerialNumber,
		Characteristics:  d.Characteristics,
		Category:         d.Category,
		Units:            d.Units,
		PowerConsumption: d.PowerConsumption,
		CreatedDate:      d.CreatedDate,
		UpdatedDate:      d.UpdatedDate,
		DeletedDate:      d.DeletedDate,
	}
}
