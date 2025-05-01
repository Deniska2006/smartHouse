package middlewares

import (
	"errors"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
)

func IsOwner() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			user := r.Context().Value(controllers.UserKey).(domain.User)
			house := r.Context().Value(controllers.HouseKey).(domain.House)

			if house.UserId != user.Id {
				err := errors.New("Acces denied")
				controllers.Forbidden(w, err)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}
