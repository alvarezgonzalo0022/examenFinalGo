package routes

import (
	"database/sql"
	"github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/handler/ping"
	handlerOdontologo "github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/handler/odontologos"
	handlerPaciente "github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/handler/pacientes"
	odontologo "github.com/alvarezgonzalo0022/examenFinalGo/internal/dentists"
	paciente "github.com/alvarezgonzalo0022/examenFinalGo/internal/patients"
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
	r.buildOdontologoRoutes()
	r.buildPacienteRoutes()
	r.buildPingRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildOdontologoRoutes maps all routes for the odontologos domain.
func (r *router) buildOdontologoRoutes() {
	// Create a new odontologos controller.
	repository := odontologo.NewMySqlRepository(r.db)
	service := odontologo.NewServiceDentist(repository)
	controlador := handlerOdontologo.NewControladorOdontologo(service)

	grupoOdontologo := r.routerGroup.Group("/odontologos")
	{
		grupoOdontologo.POST("", middleware.Authenticate(), controlador.HandlerCreate())
		grupoOdontologo.GET("", middleware.Authenticate(), controlador.HandlerGetAll())
		grupoOdontologo.GET("/:id", controlador.HandlerGetByID())
		grupoOdontologo.PUT("/:id", middleware.Authenticate(), controlador.HandlerUpdate())
		grupoOdontologo.DELETE("/:id", middleware.Authenticate(), controlador.HandlerDelete())
		grupoOdontologo.PATCH("/:id", middleware.Authenticate(), controlador.HandlerPatch())
	}

}

// buildPacienteRoutes maps all routes for the pacientes domain.
func (r *router) buildPacienteRoutes() {
	// Create a new pacientes controller.
	repository := paciente.NewPacientesMySqlRepository(r.db)
	service := paciente.NewServicePacientes(repository)
	controlador := handlerPaciente.NewControladorPaciente(service)

	grupoOdontologo := r.routerGroup.Group("/pacientes")
	{
		grupoOdontologo.POST("", middleware.Authenticate(), controlador.HandlerCreate())
		grupoOdontologo.GET("", middleware.Authenticate(), controlador.HandlerGetAll())
		grupoOdontologo.GET("/:id", controlador.HandlerGetByID())
		grupoOdontologo.PUT("/:id", middleware.Authenticate(), controlador.HandlerUpdate())
		grupoOdontologo.DELETE("/:id", middleware.Authenticate(), controlador.HandlerDelete())
		grupoOdontologo.PATCH("/:id", middleware.Authenticate(), controlador.HandlerPatch())
	}

}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControllerPing()
	r.routerGroup.GET("/ping", pingController.HandlerPing())

}
