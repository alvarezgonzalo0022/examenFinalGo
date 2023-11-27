package turnos

import (
	"context"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type RepositoryTurnos interface {
	Create(ctx context.Context, turno domain.Turno) (domain.Turno, error)
	GetAll(ctx context.Context) ([]domain.Turno, error)
	GetByID(ctx context.Context, id int) (domain.Turno, error)
	Update(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error)
}
