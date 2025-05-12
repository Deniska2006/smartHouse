package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
)

type MeasurementController struct {
	measService app.MeasurementService
}

func NewMeasurementController(ms app.MeasurementService) MeasurementController {
	return MeasurementController{
		measService: ms,
	}
}

func (c MeasurementController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		meas, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		err = c.measService.FindDeviceByUUID(meas.DeviceUUID)

		if err != nil {
			log.Printf("MeasurementController.FindDeviceByUUID(c.measService:FindDeviceByUUID): %s", err)
			BadRequest(w, errors.New("Device doesnt exist or acces denied"))
			return
		}

		err = c.measService.Save(meas)
		if err != nil {
			log.Printf("MeasurementController.Save(c.measService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

	}
}
