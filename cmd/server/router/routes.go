package routes

import (
	"database/sql"

	handlerAppointment "github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/handler/appointments"
	handlerDentist "github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/handler/dentists"
	handlerPatient "github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/handler/patients"
	"github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/handler/ping"
	appointments "github.com/alvarezgonzalo0022/examenFinalGo/internal/appointments"
	dentists "github.com/alvarezgonzalo0022/examenFinalGo/internal/dentists"
	patients "github.com/alvarezgonzalo0022/examenFinalGo/internal/patients"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// Router interface defines the methods that any router must implement.
type Router interface {
	MapRoutes()
}

// router is the Gin router.
type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

// NewRouter creates a new Gin router.
func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

// MapRoutes maps all routes.
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildDentistRoutes()
	r.buildPatientRoutes()
	r.buildAppointmentRoutes()
	r.buildPingRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildOdontologoRoutes maps all routes for the odontologos domain.
func (r *router) buildDentistRoutes() {
	// Create a new odontologos controller.
	repository := dentists.NewMySqlRepository(r.db)
	service := dentists.NewServiceDentist(repository)
	controller := handlerDentist.NewControllerDentist(service)

	groupDentists := r.routerGroup.Group("/dentists")
	{
		groupDentists.POST("", middleware.Authenticate(), controller.HandlerCreate())
		groupDentists.GET("", middleware.Authenticate(), controller.HandlerGetAll())
		groupDentists.GET("/:id", controller.HandlerGetByID())
		groupDentists.PUT("/:id", middleware.Authenticate(), controller.HandlerUpdate())
		groupDentists.DELETE("/:id", middleware.Authenticate(), controller.HandlerDelete())
		groupDentists.PATCH("/:id", middleware.Authenticate(), controller.HandlerPatch())
	}

}

// build PatientRoutes maps all routes for the patients domain.
func (r *router) buildPatientRoutes() {
	// Create a new patients controller.
	repository := patients.NewPatientsMySqlRepository(r.db)
	service := patients.NewServicePatients(repository)
	controller := handlerPatient.NewControllerPatient(service)

	groupPatients := r.routerGroup.Group("/patients")
	{
		groupPatients.POST("", middleware.Authenticate(), controller.HandlerCreate())
		groupPatients.GET("", middleware.Authenticate(), controller.HandlerGetAll())
		groupPatients.GET("/:id", controller.HandlerGetByID())
		groupPatients.PUT("/:id", middleware.Authenticate(), controller.HandlerUpdate())
		groupPatients.DELETE("/:id", middleware.Authenticate(), controller.HandlerDelete())
		groupPatients.PATCH("/:id", middleware.Authenticate(), controller.HandlerPatch())
	}
}

func (r *router) buildAppointmentRoutes() {
	// Create a new patients controller.
	repository := appointments.NewMySqlRepository(r.db)
	service := appointments.NewServiceAppointments(repository)
	controller := handlerAppointment.NewControllerAppointment(service)

	groupAppointments := r.routerGroup.Group("/appointments")
	{
		// groupAppointments.POST("", middleware.Authenticate(), controller.HandlerCreate())
		groupAppointments.GET("", middleware.Authenticate(), controller.HandlerGetAll())
		groupAppointments.GET("/:id", controller.HandlerGetByID())
		// groupAppointments.PUT("/:id", middleware.Authenticate(), controller.HandlerUpdate())
		 groupAppointments.DELETE("/:id", middleware.Authenticate(), controller.HandlerDelete())
		// groupAppointments.PATCH("/:id", middleware.Authenticate(), controller.HandlerPatch())
	}
}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControllerPing()
	r.routerGroup.GET("/ping", pingController.HandlerPing())

}
