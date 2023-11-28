package turnos

import (
	"context"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/turnos"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
	"github.com/gin-gonic/gin")

	type turnoHandler struct {
		service turnos.ServiceTurnos
	}
		
	func NewTurnoHandler(service turnos.ServiceTurnos) *turnoHandler {
		return &turnoHandler{
			service: service,
		}
	}	
 
// Paciente godoc
//	@Summary		turno example
//	@Description	Create a new turno
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/turnos [post]


func (h *turnoHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurno

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		response, err := c.service.Create(ctx, request)
	
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s. %s", "Internal Server Error", err)
			return
		}
		
		web.Success(ctx, http.StatusCreated, response)
	}
}



	
