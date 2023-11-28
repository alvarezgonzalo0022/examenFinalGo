package appointments

import (
	"context"
	"log"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type ServiceAppointments interface {
	Create(ctx context.Context, appointment domain.Appointment) (domain.Appointment, error)
	GetAll(ctx context.Context) ([]domain.Appointment, error)
	GetByID(ctx context.Context, id int) (domain.Appointment, error)
	Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error)
}

type service struct {
	repository RepositoryAppointments
}

func NewServiceAppointments(repository RepositoryAppointments) ServiceAppointments {
	return &service{repository: repository}
}

// Create is a method that creates a new appointment.
func (s *service) Create(ctx context.Context, appointment domain.Appointment) (domain.Appointment, error) {
	appointment, err := s.repository.Create(ctx, appointment)
	if err != nil {
		log.Println("[AppointmentsService][Create] error creating appointment", err)
		return domain.Appointment{}, err
	} 
    return appointment, nil;
}

// GetAll is a method that returns all appointment.
func (s *service) GetAll(ctx context.Context) ([]domain.Appointment, error) {
	appointments, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[AppointmentsService][GetAll] error getting all appointments", err)
		return []domain.Appointment{}, err
	}

	return appointments, nil
}

// GetByID is a method that returns a appointment by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Appointment, error) {
	appointment, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[AppointmentsService][GetByID] error getting appointment by ID", err)
		return domain.Appointment{}, err
	}

	return appointment, nil
}

// Update is a method that updates a appointment by ID.
func (s *service) Update(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error) {
	appointment, err := s.repository.Update(ctx, appointment, id)
	if err != nil {
		log.Println("[AppointmentsService][Update] error updating appointment by ID", err)
		return domain.Appointment{}, err
	}

	return appointment, nil
}

// Delete is a method that deletes a appointment by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[AppointmentsService][Delete] error deleting appointment by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a appointment by ID.
func (s *service) Patch(ctx context.Context, appointment domain.Appointment, id int) (domain.Appointment, error) {
	appointment, err := s.repository.Patch(ctx, appointment, id)
	if err != nil {
		log.Println("[AppointmentsService][Patch] error patching appointment by ID", err)
		return domain.Appointment{}, err
	}

	return appointment, nil
}

func (s *service) validatePatch(appointmentStore domain.Appointment, appointment domain.Appointment) (domain.Appointment, error) {
	if appointmentStore.Id != appointment.Id {
		return domain.Appointment{}, domain.ErrInvalidID
	}

	if appointmentStore.AppointmentDate != appointment.AppointmentDate {
		appointmentStore.AppointmentDate = appointment.AppointmentDate
	}
	
	if appointmentStore.AppointmentTime != appointment.AppointmentTime {
		appointmentStore.AppointmentTime = appointment.AppointmentTime
	}

	if appointmentStore.IdPatient != appointment.IdPatient {
		appointmentStore.IdPatient = appointment.IdPatient
	}

	if appointmentStore.IdDentist != appointment.IdDentist {
		appointmentStore.IdDentist = appointment.IdDentist
	}

	if appointmentStore.Description != appointment.Description {
		appointmentStore.Description = appointment.Description
	}

	return appointmentStore, nil
}

