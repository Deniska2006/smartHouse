package container

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/config"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/middlewares"
	"github.com/go-chi/jwtauth/v5"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type Container struct {
	Middlewares
	Services
	Controllers
}

type Middlewares struct {
	AuthMw func(http.Handler) http.Handler
}

type Services struct {
	app.AuthService
	app.UserService
	app.HouseService
	app.RoomService
	app.DeviceService
	app.MeasurementService
	app.EventService
}

type Controllers struct {
	AuthController        controllers.AuthController
	UserController        controllers.UserController
	HouseController       controllers.HouseController
	RoomController        controllers.RoomController
	DeviceController      controllers.DeviceController
	MeasurementController controllers.MeasurementController
	EventController       controllers.EventController
}

func New(conf config.Configuration) Container {
	tknAuth := jwtauth.New("HS256", []byte(conf.JwtSecret), nil)
	sess := getDbSess(conf)

	sessionRepository := database.NewSessRepository(sess)
	userRepository := database.NewUserRepository(sess)
	houseRepository := database.NewHouseRepository(sess)
	roomRepository := database.NewRoomRepository(sess)
	deviceRepository := database.NewDeviceRepository(sess)
	measurementRepositry := database.NewMeasurementRepository(sess)
	eventRepository := database.NewEventRepository(sess)

	userService := app.NewUserService(userRepository)
	authService := app.NewAuthService(sessionRepository, userRepository, tknAuth, conf.JwtTTL)
	houseService := app.NewHouseService(houseRepository, roomRepository)
	roomService := app.NewRoomService(roomRepository, deviceRepository)
	deviceService := app.NewDeviceService(deviceRepository)
	measurementService := app.NewMeasurementService(measurementRepositry, deviceRepository)
	eventService := app.NewEventtService(eventRepository, deviceRepository)

	authController := controllers.NewAuthController(authService, userService)
	userController := controllers.NewUserController(userService, authService)
	houseController := controllers.NewHouseController(houseService)
	roomController := controllers.NewRoomController(roomService)
	deviceController := controllers.NewDeviceController(deviceService)
	measurementController := controllers.NewMeasurementController(measurementService)
	eventController := controllers.NewEventController(eventService)

	authMiddleware := middlewares.AuthMiddleware(tknAuth, authService, userService)

	return Container{
		Middlewares: Middlewares{
			AuthMw: authMiddleware,
		},
		Services: Services{
			authService,
			userService,
			houseService,
			roomService,
			deviceService,
			measurementService,
			eventService,
		},
		Controllers: Controllers{
			authController,
			userController,
			houseController,
			roomController,
			deviceController,
			measurementController,
			eventController,
		},
	}
}

func getDbSess(conf config.Configuration) db.Session {
	sess, err := postgresql.Open(
		postgresql.ConnectionURL{
			User:     conf.DatabaseUser,
			Host:     conf.DatabaseHost,
			Password: conf.DatabasePassword,
			Database: conf.DatabaseName,
		})
	if err != nil {
		log.Fatalf("Unable to create new DB session: %q\n", err)
	}
	return sess
}
