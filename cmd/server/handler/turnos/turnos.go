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
 
	func (h *turnoHandler) Create(c *gin.Context) {
		var turno domain.Turno
		if err := c.ShouldBindJSON(&turno); err != nil {
			web.NewResponse(c).BadRequest(err)
			return
		}
		turno, err := h.service.Create(context.Background(), turno)
		if err != nil {
			web.NewResponse(c).InternalServerError(err)
			return
		}
		web.NewResponse(c).Ok(turno)
	}

	
