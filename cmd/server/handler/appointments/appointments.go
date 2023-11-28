package appointments

import (
	"net/http"

	// "github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/appointments"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
	"github.com/gin-gonic/gin"
)

	type Controller struct {
		service appointments.ServiceAppointments
	}
		
	func NewControllerAppointment(service appointments.ServiceAppointments) *Controller {
		return &Controller{
			service: service,
		}
	}	

	func (c *Controller) HandlerGetAll() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			appointments, err := c.service.GetAll(ctx)
	
			if err != nil {
				web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
				return
			}
	
			web.Success(ctx, http.StatusOK, gin.H{
				"data": appointments,
			})
		}
	}

	
