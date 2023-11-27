package turnos

import (
	"context"
	"log"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type ServiceTurnos interface {
	Create(ctx context.Context, turno domain.Turno) (domain.Turno, error)
	GetAll(ctx context.Context) ([]domain.Turno, error)
	GetByID(ctx context.Context, id int) (domain.Turno, error)
	Update(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error)
}

type service struct {
	repository RepositoryTurnos
}

func NewServiceTurnos(repository RepositoryTurnos) ServiceTurnos {
	return &service{repository: repository}
}

// Create is a method that creates a new turno.
func (s *service) Create(ctx context.Context, turno domain.Turno) (domain.Turno, error) {
	turno, err := s.repository.Create(ctx, turno)
	if err != nil {
		log.Println("[TurnosService][Create] error creating turno", err)
		return domain.Turno{}, err
	} 
     return turno, nil;
}

// GetAll is a method that returns all turno.
func (s *service) GetAll(ctx context.Context) ([]domain.Turno, error) {
	listTurnos, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[TurnosService][GetAll] error getting all turnos", err)
		return []domain.Turno{}, err
	}

	return listTurnos, nil
}

// GetByID is a method that returns a turno by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Turno, error) {
	turno, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[TurnosService][GetByID] error getting turno by ID", err)
		return domain.Turno{}, err
	}

	return turno, nil
}

// Update is a method that updates a turno by ID.
func (s *service) Update(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error) {
	turno, err := s.repository.Update(ctx, turno, id)
	if err != nil {
		log.Println("[TurnosService][Update] error updating turno by ID", err)
		return domain.Turno{}, err
	}

	return turno, nil
}

// Delete is a method that deletes a turno by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[TurnosService][Delete] error deleting turno by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a turno by ID.
func (s *service) Patch(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error) {
	turno, err := s.repository.Patch(ctx, turno, id)
	if err != nil {
		log.Println("[TurnosService][Patch] error patching turno by ID", err)
		return domain.Turno{}, err
	}

	return turno, nil
}

func (s *service) validatePatch(turnoStore domain.Turno, turno domain.Turno) (domain.Turno, error) {
	if turnoStore.Id != turno.Id {
		return domain.Turno{}, domain.ErrInvalidID
	}

	if turnoStore.AppointmentDate != turno.AppointmentDate {
		turnoStore.AppointmentDate = turno.AppointmentDate
	}
	
	if turnoStore.AppointmentTime != turno.AppointmentTime {
		turnoStore.AppointmentTime = turno.AppointmentTime
	}

	if turnoStore.IdPatient != turno.IdPatient {
		turnoStore.IdPatient = turno.IdPatient
	}

	if turnoStore.IdDentist != turno.IdDentist {
		turnoStore.IdDentist = turno.IdDentist
	}

	if turnoStore.Description != turno.Description {
		turnoStore.Description = turno.Description
	}

	return turnoStore, nil
}

