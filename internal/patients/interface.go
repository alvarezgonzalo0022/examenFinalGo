package patients

import (
	"context"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type RepositoryPatients interface {
	Create(ctx context.Context, patient domain.Patient) (domain.Patient, error)
	GetAll(ctx context.Context) ([]domain.Patient, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
}
