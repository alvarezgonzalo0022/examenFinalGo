package patients

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	patients "github.com/alvarezgonzalo0022/examenFinalGo/internal/patients"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service patients.ServicePatients
}

func NewControllerPatient(service patients.ServicePatients) *Controller {
	return &Controller{
		service: service,
	}
}

// Patient godoc
// @Summary patient example
// @Description Create a new patient
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patients [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Patient

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		patient, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patient)

	}
}

// Patient godoc
// @Summary patient example
// @Description Get all patients
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 500 {object} web.errorResponse
// @Router /patients [get]
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		patients, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patients)
	}
}

// Patient godoc
// @Summary patient example
// @Description Get patient by id
// @Tags patient
// @Param id path int true "id del patient"
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patients/:id [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Printf("Error al convertir ID: %v", err)
			web.Error(ctx, http.StatusBadRequest, "%s", "invalid id")
			return
		}

		patient, err := c.service.GetByID(ctx, id)
		if err != nil {
			log.Printf("Error al obtener patient por ID: %v", err)
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patient)
	}
}

// Patient godoc
// @Summary patient example
// @Description Update patient by id
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patients/:id [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Patient

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		patient, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patient)

	}
}

// Patient godoc
// @Summary patient example
// @Description Delete patient by id
// @Tags patient
// @Param id path int true "id del patient"
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patients/:id [delete]
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
			"mensaje": "deleted patient",
		})
	}
}

// Patient godoc
// @Summary patient example
// @Description Patch patient
// @Tags patient
// @Param id path int true "id del patient"
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patients/:id [patch]
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "invalid id")
			return
		}

		var request domain.Patient

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		patient, err := c.service.Patch(ctx, request, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, patient)
	}
}
