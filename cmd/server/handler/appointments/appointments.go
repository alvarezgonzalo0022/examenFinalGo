package appointments

import (
	"log"
	"net/http"
	"strconv"

	// "github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/appointments"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/dentists"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/patients"

	// "github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service 		appointments.ServiceAppointments
	patientService 	patients.ServicePatients
	dentistService 	dentists.ServiceDentists
}

func NewControllerAppointment(service appointments.ServiceAppointments, patientService patients.ServicePatients,
    dentistService dentists.ServiceDentists,) *Controller {
	return &Controller{
		service: service,
		patientService: patientService,
		dentistService: dentistService,
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

		_, err = c.dentistService.GetByID(ctx, request.DentistId)
        if err != nil {
            log.Println("Error getting dentist:", err)
            web.Error(ctx, http.StatusBadRequest, "%s", "invalid dentist")
            return
        }

        // Validar la existencia del paciente
        _, err = c.patientService.GetByID(ctx, request.PatientId)
        if err != nil {
            log.Println("Error getting patient:", err)
            web.Error(ctx, http.StatusBadRequest, "%s", "invalid patient")
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
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error", err)
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
// @Success 200 {object}
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
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error" , err)
			return
		}

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