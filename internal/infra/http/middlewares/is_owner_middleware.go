package middlewares

import (
	"errors"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
)

func IsOwner() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(controllers.UserKey).(domain.User)
			house := r.Context().Value(controllers.HouseKey).(domain.House)
			if house.UserId != user.Id {
				controllers.Forbidden(w, errors.New("Access denied: user doesn't own the house"))
				return
			}


			roomCtx := r.Context().Value(controllers.RoomKey)
			if roomCtx != nil {
				room := roomCtx.(domain.Room)
				if room.HouseId != house.Id {
					controllers.Forbidden(w, errors.New("Access denied: room doesn't belong to house"))
					return
				}
			}

			deviceCtx := r.Context().Value(controllers.DeviceKey)
			if deviceCtx != nil {
				device := deviceCtx.(domain.Device)
				room := roomCtx.(domain.Room)
				if device.RoomId != room.Id {
					controllers.Forbidden(w, errors.New("Access denied: room doesn't have this device"))
					return
				}
			}

			
			next.ServeHTTP(w, r)
		})
	}
}

