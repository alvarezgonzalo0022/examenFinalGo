package appointments

import (
	"log"
	"net/http"
	"strconv"

	// "github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/appointments"
	// "github.com/alvarezgonzalo0022/examenFinalGo/internal/dentists"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	// "github.com/alvarezgonzalo0022/examenFinalGo/internal/patients"

	// "github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service 		appointments.ServiceAppointments
}

func NewControllerAppointment(service appointments.ServiceAppointments,) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.AppointmentRequest

		err := ctx.Bind(&request)
		if err != nil {
			log.Println("Error binding request:", err)
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		appointment, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, appointment)

	}
}

func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appointments, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, appointments)
	}
}

/* --------------------------------- GET BY ID ------------------------------- */
// Appointment godoc
// @Summary appointment example
// @Description Get appointment by id
// @Tags appointment
// @Param id path int true "id del appointment"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointments/{id} [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "invalid id")
			return
		}

		appointment, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, appointment)
	}
}

// Appointment godoc
// @Summary Update appointment by ID
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointments/:id [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ... (manejo de errores, validaciones, etc.)

		// Obtener el ID de la URL
		id := ctx.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		// Crear una estructura de dominio para la actualización
		var request domain.AppointmentRequest
		errBind := ctx.Bind(&request)
		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		// Actualizar la cita
		appointment, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			// Manejar el error (por ejemplo, enviar una respuesta de error al cliente)
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		// Responder con la cita actualizada
		web.Success(ctx, http.StatusOK, appointment)
	}
}



/* --------------------------------- DELETE ------------------------------- */
// Appointment godoc
// @Summary appointment example
// @Description Delete appointment by id
// @Tags appointment
// @Param id path int true "id del appointment"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "invalid id")
			return
		}

		_, err = c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "%s", "appointment not found")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"mensaje": "deleted appointment",
		})
	}
}

/* --------------------------------- PATCH ------------------------------- */
// Appointment godoc
// @Summary appointment example
// @Description Patch appointment by id
// @Tags appointment
// @Param id path int true "id del appointment"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [patch]
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ... (manejo de errores, validaciones, etc.)

		// Obtener el ID de la URL
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "invalid id")
			return
		}

		

		// Crear una estructura de dominio para el parche
		var request domain.AppointmentPatchRequest
		errBind := ctx.Bind(&request)
		if errBind != nil {
			log.Println("Error binding request:", errBind)
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		// Aplicar el parche a la cita
		appointment, err := c.service.Patch(ctx, request, id)
		if err != nil {
			// Manejar el error (por ejemplo, enviar una respuesta de error al cliente)
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		// Responder con la cita actualizada después del parche
		web.Success(ctx, http.StatusOK, appointment)
	}
}