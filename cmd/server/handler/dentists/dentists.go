package dentists

import (
	"net/http"
	"strconv"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	odontologos "github.com/alvarezgonzalo0022/examenFinalGo/internal/dentists"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service odontologos.ServiceDentists
}

func NewControllerDentist(service odontologos.ServiceDentists) *Controller {
	return &Controller{
		service: service,
	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description Create a new odontologo
// @Tags odontologo
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Dentist

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		dentist, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentist)

	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description Get all odontologos
// @Tags odontologo
// @Accept json
// @Produce json
// @Success 200 {object} 
// @Failure 500 {object} web.errorResponse
// @Router /odontologos [get]
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentists, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentists)
	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description Get odontologo by id
// @Tags odontologo
// @Param id path int true "id del odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} 
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "invalid id")
			return
		}

		dentist, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentist)
	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description Update odontologo by id
// @Tags odontologo
// @Accept json
// @Produce json
// @Success 200 {object} 
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Dentist

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

		dentist, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentist)

	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description Delete odontologo by id
// @Tags odontologo
// @Param id path int true "id del odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} 
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"message": "deleted dentist",
		})
	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description Patch odontologo
// @Tags odontologo
// @Param id path int true "id del odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} 
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [patch]
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "invalid id")
			return
		}

		var request domain.Dentist

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		dentist, err := c.service.Patch(ctx, request, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentist)
	}
}
