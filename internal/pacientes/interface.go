package pacientes

import (
	"context"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type RepositoryPacientes interface {
	Create(ctx context.Context, paciente domain.Paciente) (domain.Paciente, error)
	GetAll(ctx context.Context) ([]domain.Paciente, error)
	GetByID(ctx context.Context, id int) (domain.Paciente, error)
	Update(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, paciente domain.Paciente, id int) (domain.Paciente, error)
}
