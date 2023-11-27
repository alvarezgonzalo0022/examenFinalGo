package odontologos

import (
	"context"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type RepositoryOdontologos interface {
	Create(ctx context.Context, odontologo domain.Odontologo) (domain.Odontologo, error)
	GetAll(ctx context.Context) ([]domain.Odontologo, error)
	GetByID(ctx context.Context, id int) (domain.Odontologo, error)
	Update(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, odontologo domain.Odontologo, id int) (domain.Odontologo, error)
}
