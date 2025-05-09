package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/BohdanBoriak/boilerplate-go-back/config"
	"github.com/BohdanBoriak/boilerplate-go-back/config/container"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5/middleware"
)

func Router(cont container.Container) http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.RedirectSlashes, middleware.Logger, cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "capacitor://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/api", func(apiRouter chi.Router) {
		// Health
		apiRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())
			healthRouter.Handle("/*", NotFoundJSON())
		})

		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			// Public routes
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Route("/auth", func(apiRouter chi.Router) {
					AuthRouter(apiRouter, cont.AuthController, cont.AuthMw)
				})
			})

			// Protected routes
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Use(cont.AuthMw)

				UserRouter(apiRouter, cont.UserController)
				HouseRouter(apiRouter, cont.HouseController, cont.HouseService)
				RoomRouter(apiRouter, cont.RoomController, cont.HouseService, cont.RoomService)
				DeviceRouter(apiRouter, cont.DeviceController, cont.HouseService, cont.RoomService, cont.DeviceService)

				apiRouter.Handle("/*", NotFoundJSON())
			})
		})
	})

	router.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		workDir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(workDir, config.GetConfiguration().FileStorageLocation))
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})

	return router
}

func AuthRouter(r chi.Router, ac controllers.AuthController, amw func(http.Handler) http.Handler) {
	r.Route("/", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/register",
			ac.Register(),
		)
		apiRouter.Post(
			"/login",
			ac.Login(),
		)
		apiRouter.With(amw).Post(
			"/logout",
			ac.Logout(),
		)
	})
}

func UserRouter(r chi.Router, uc controllers.UserController) {
	r.Route("/users", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			uc.FindMe(),
		)
		apiRouter.Put(
			"/",
			uc.Update(),
		)
		apiRouter.Delete(
			"/",
			uc.Delete(),
		)
	})
}

func HouseRouter(r chi.Router, hc controllers.HouseController, hs app.HouseService) {
	hpom := middlewares.PathObject("houseId", controllers.HouseKey, hs)
	r.Route("/houses", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/",
			hc.Save(),
		)
		apiRouter.With(hpom, middlewares.IsOwner()).Get( // middlewares.IsOwner()
			"/{houseId}",
			hc.Find(),
		)
		apiRouter.With().Get(
			"/",
			hc.FindList(),
		)
		apiRouter.With(hpom, middlewares.IsOwner()).Put(
			"/{houseId}",
			hc.Update(),
		)
		apiRouter.With(hpom, middlewares.IsOwner()).Delete(
			"/{houseId}",
			hc.Delete(),
		)
	})
}

func RoomRouter(r chi.Router, rc controllers.RoomController, hs app.HouseService, rs app.RoomService) {
	hpom := middlewares.PathObject("houseId", controllers.HouseKey, hs)
	rpom := middlewares.PathObject("roomId", controllers.RoomKey, rs)

	r.Route("/houses/{houseId}/rooms", func(apiRouter chi.Router) {
		apiRouter.With(hpom).Post(
			"/",
			rc.Save(),
		)
		apiRouter.With(hpom, middlewares.IsOwner()).Get(
			"/",
			rc.FindList(),
		)
		apiRouter.With(hpom, rpom, middlewares.IsOwner()).Get(
			"/{roomId}",
			rc.Find(),
		)
		apiRouter.With(hpom, rpom, middlewares.IsOwner()).Put(
			"/{roomId}",
			rc.Update(),
		)
		apiRouter.With(hpom, rpom, middlewares.IsOwner()).Delete(
			"/{roomId}",
			rc.Delete(),
		)

	})
}

func DeviceRouter(r chi.Router, dc controllers.DeviceController, hs app.HouseService, rs app.RoomService, ds app.DeviceService) {
	hpom := middlewares.PathObject("houseId", controllers.HouseKey, hs)
	rpom := middlewares.PathObject("roomId", controllers.RoomKey, rs)
	dpom := middlewares.PathObject("deviceId", controllers.DeviceKey, ds)

	r.Route("/houses/{houseId}/rooms/{roomId}/devices", func(apiRouter chi.Router) {
		apiRouter.With(hpom, rpom).Post(
			"/",
			dc.Save(),
		)
		apiRouter.With(hpom, rpom, middlewares.IsOwner()).Get(
			"/",
			dc.FindList(),
		)
		apiRouter.With(hpom, rpom, dpom, middlewares.IsOwner()).Get(
			"/{deviceId}",
			dc.Find(),
		)
		apiRouter.With(hpom, rpom, dpom, middlewares.IsOwner()).Put(
			"/{deviceId}",
			dc.Update(),
		)
	})

}

func NotFoundJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode("Resource Not Found")
		if err != nil {
			fmt.Printf("writing response: %s", err)
		}
	}
}

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("Ok")
		if err != nil {
			fmt.Printf("writing response: %s", err)
		}
	}
}
