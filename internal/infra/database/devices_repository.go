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
	FindList(rId uint64) ([]domain.Device, error)
	Find(id uint64) (domain.Device, error)
	Update(updt domain.Device, d domain.Device) (domain.Device, error)
	Delete(dId uint64) error
	FindDeviceByUUID(u uuid.UUID) error
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

func (r deviceRepository) FindList(rId uint64) ([]domain.Device, error) {
	var devices []device
	err := r.coll.
		Find(db.Cond{
			"room_id":      rId,
			"deleted_date": nil}).All(&devices)
	if err != nil {
		return nil, err
	}

	return r.mapModelToDomainCollection(devices), nil
}

func (r deviceRepository) Find(id uint64) (domain.Device, error) {
	var d device
	err := r.coll.
		Find(db.Cond{
			"id":           id,
			"deleted_date": nil}).One(&d)
	if err != nil {
		return domain.Device{}, err
	}

	return r.mapModelToDomain(d), nil
}

func (r deviceRepository) Update(updt domain.Device, d domain.Device) (domain.Device, error) {
	err := r.coll.Find(db.Cond{"id": d.Id}).Update(r.mapDomainToInterfaceUpdate(updt, d))
	if err != nil {
		return domain.Device{}, err
	}

	return r.mapUpdateToDomain(updt, d), nil
}

func (r deviceRepository) Delete(dId uint64) error {
	err := r.coll.Find(db.Cond{"id": dId}).Update(map[string]interface{}{
		"deleted_date": time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (r deviceRepository) FindDeviceByUUID(u uuid.UUID) error {
	var d domain.Device
	err := r.coll.Find(db.Cond{"uuid": u}).One(&d)
	if err == db.ErrNoMoreRows {
		return err // не знайдено
	}
	if err != nil {
		return err
	}
	return nil
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

func (r deviceRepository) mapModelToDomainCollection(devices []device) []domain.Device {
	dvcs := make([]domain.Device, len(devices))
	for i, d := range devices {
		dvcs[i] = r.mapModelToDomain(d)
	}

	return dvcs
}

func (r deviceRepository) mapDomainToInterfaceUpdate(updt, d domain.Device) map[string]interface{} {
	result := make(map[string]interface{})

	if updt.SerialNumber != "" {
		result["serial_number"] = updt.SerialNumber
	}
	if updt.Characteristics != nil {
		result["characteristics"] = updt.Characteristics
	}
	if updt.Category != "" {
		result["category"] = updt.Category
	}

	finalCategory := d.Category
	if updt.Category != "" {
		finalCategory = updt.Category
	}

	if finalCategory == domain.SENSOR {
		result["power_consumption"] = nil
		if updt.Units != nil {
			result["units"] = updt.Units
		}
	}
	if finalCategory == domain.ACTUATOR {
		result["units"] = nil
		if updt.PowerConsumption != nil {
			result["power_consumption"] = updt.PowerConsumption
		}
	}

	result["updated_date"] = time.Now()
	return result
}

func (r deviceRepository) mapUpdateToDomain(updt domain.Device, d domain.Device) domain.Device {
	if updt.SerialNumber != "" {
		d.SerialNumber = updt.SerialNumber
	}
	if updt.Characteristics != nil {
		d.Characteristics = updt.Characteristics
	}
	if updt.Category != "" {
		d.Category = updt.Category
	}
	if updt.Units != nil {
		d.Units = updt.Units
	}
	if updt.PowerConsumption != nil {
		d.PowerConsumption = updt.PowerConsumption
	}
	if updt.Category == domain.ACTUATOR {
		d.Units = nil
	}
	if updt.Category == domain.SENSOR {
		d.PowerConsumption = nil
	}
	d.UpdatedDate = time.Now()
	return d

}
