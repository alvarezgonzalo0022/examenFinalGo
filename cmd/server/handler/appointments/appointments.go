package appointments

import (
	"log"
	"net/http"
	"strconv"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/appointments"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
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

// godoc Appointments
// HandlerCreate crea una nueva cita.
// @Summary Crear cita
// @Description Crea una nueva cita.
// @Tags appointments
// @Accept json
// @Produce json
// Param request body domain.AppointmentRequest true "Información de la cita a crear"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /appointments [post]
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

// godoc Appointments
// HandlerGetAll obtiene todas las citas.
// @Summary Obtener todas las citas
// @Description Obtiene todas las citas.
// @Tags appointments
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 500 {object} web.ErrorResponse
// @Router /appointments [get]
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
// godoc Appointments
// HandlerGetByID obtiene una cita por su ID.
// @Summary Obtener cita por ID
// @Description Obtiene una cita por su ID.
// @Tags appointments
// @Param id path int true "ID de la cita"
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
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

// godoc Appointments
// HandlerUpdate actualiza una cita por su ID.
// @Summary Actualizar cita por ID
// @Description Actualiza una cita por su ID.
// @Tags appointments
// @Param id path int true "ID de la cita"
// @Accept json
// @Produce json
// Param request body domain.AppointmentRequest true "Información de la cita a actualizar"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /appointments/{id} [put]
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
// godoc Appointments
// HandlerDelete elimina una cita por su ID.
// @Summary Eliminar cita por ID
// @Description Elimina una cita por su ID.
// @Tags appointments
// @Param id path int true "ID de la cita"
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /appointments/{id} [delete]
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
// godoc Appointments
// HandlerPatch aplica un parche a una cita por su ID.
// @Summary Aplicar parche a cita por ID
// @Description Aplica un parche a una cita por su ID.
// @Tags appointments
// @Param id path int true "ID de la cita"
// @Accept json
// @Produce json
// @Param request body domain.AppointmentPatchRequest true "Parche a aplicar a la cita"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.ErrorResponse
// @Failure 404 {object} web.ErrorResponse
// @Failure 500 {object} web.ErrorResponse
// @Router /appointments/{id} [patch]
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