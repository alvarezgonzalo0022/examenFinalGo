package turnos

import (
	"net/http"

	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/turnos"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
	"github.com/gin-gonic/gin"
)

	type AppointmentController struct {
		service turnos.ServiceTurnos
	}
		
	func NewControllerAppoitment(service turnos.ServiceTurnos) *AppointmentController {
		return &AppointmentController{
			service: service,
		}
	}	

	func (c *AppointmentController) HandlerGetAll() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			pacientes, err := c.service.GetAll(ctx)
	
			if err != nil {
				web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
				return
			}
	
			web.Success(ctx, http.StatusOK, gin.H{
				"data": pacientes,
			})
		}
	}

	
