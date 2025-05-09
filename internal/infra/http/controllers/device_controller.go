package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"github.com/google/uuid"
)

type DeviceController struct {
	deviceService app.DeviceService
}

func NewDeviceController(dvc app.DeviceService) DeviceController {
	return DeviceController{
		deviceService: dvc,
	}
}

func (c DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		if (device.PowerConsumption != nil && device.Units != nil) || // не може бути і те, і те
			(device.Category != domain.SENSOR && device.Category != domain.ACTUATOR) || // має бути одна з 2 категорій а не щось третє
			(device.Category == domain.SENSOR && device.Units == nil) || // SENSOR → Units обов'язково
			(device.Category == domain.ACTUATOR && device.PowerConsumption == nil) { // ACTUATOR → PowerConsumption обов'язково

			err = errors.New("Wrong data request")
			Forbidden(w, err)
			return
		}

		device.HouseId = r.Context().Value(HouseKey).(domain.House).Id
		device.RoomId = r.Context().Value(RoomKey).(domain.Room).Id
		device.UUID = uuid.New()

		device, err = c.deviceService.Save(device)
		if err != nil {
			log.Printf("DeviceController.Save(c.deviceService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDto(device))
	}
}

func (c DeviceController) FindList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		devices, err := c.deviceService.FindList(r.Context().Value(RoomKey).(domain.Room).Id)
		if err != nil {
			log.Printf("DeviceController.Save(c.deviceService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDtoCollection(devices))
	}
}

func (c DeviceController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Success(w, resources.DeviceDto{}.DomainToDto(r.Context().Value(DeviceKey).(domain.Device)))
	}
}

func (c DeviceController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updt, err := requests.Bind(r, requests.UpdateDeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Update(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		device := r.Context().Value(DeviceKey).(domain.Device)

		if !UpdateValid(updt, device) {
			err = errors.New("Wrong data request")
			Forbidden(w, err)
			return
		}

		device, err = c.deviceService.Update(updt, device)
		if err != nil {
			log.Printf("DeviceController.Update(c.deviceService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDto(device))
	}
}

func (c DeviceController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := c.deviceService.Delete(r.Context().Value(DeviceKey).(domain.Device).Id)
		if err != nil {
			log.Printf("DeviceController.Delete(c.deviceService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, resources.Message{Response: "Device was deleted"})
	}
}

func UpdateValid(updt, device domain.Device) bool {
	if updt.Units != nil && updt.PowerConsumption != nil { // те і те
		return false
	}
	if updt.Category == domain.ACTUATOR && updt.Units != nil { // для актуатора тільки повер
		return false
	}
	if updt.Category == domain.SENSOR && updt.PowerConsumption != nil { // для сенсора тільки юнітс
		return false
	}
	if updt.Category == "" && device.Category == domain.ACTUATOR && updt.Units != nil { // не оновлюємо категорію але намагаємося в актуатор засунути юнітся
		return false
	}
	if updt.Category == "" && device.Category == domain.SENSOR && updt.PowerConsumption != nil { // не оновлюємо категорію але намагаємося в сенсор засунути повер
		return false
	}
	if updt.Category == domain.ACTUATOR && device.Units != nil && updt.PowerConsumption == nil { // намагаємося змінити категорію не змінивши тип
		return false
	}
	if updt.Category == domain.SENSOR && device.PowerConsumption != nil && updt.Units == nil { // намагаємося змінити категорію не змінивши тип
		return false
	}
	if updt.Category != "" && updt.Units == nil && updt.PowerConsumption == nil { // якщо змінюємо категорію але не вказуємо нічого
		return false
	}
	return true
}
