package appointments

import (
	"context"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type RepositoryAppointments interface {
	Create(ctx context.Context, appointment domain.AppointmentRequest) (domain.AppointmentRequest, error)
	GetAll(ctx context.Context) ([]domain.AppointmentResponse, error)
	GetByID(ctx context.Context, id int) (domain.AppointmentResponse, error)
	Update(ctx context.Context, appointment domain.AppointmentRequest, id int) (domain.AppointmentRequest, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, appointment domain.AppointmentPatchRequest, id int) (domain.AppointmentPatchRequest, error)
}
