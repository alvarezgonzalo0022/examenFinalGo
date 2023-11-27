package turnos

import (
	"errors"
	"log"
	"time"
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

	
