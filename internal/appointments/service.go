package appointments

import (
	"context"
	"log"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type ServiceAppointments interface {
	Create(ctx context.Context, appointment domain.AppointmentRequest) (domain.AppointmentRequest, error)
	GetAll(ctx context.Context) ([]domain.AppointmentResponse, error)
	GetByID(ctx context.Context, id int) (domain.AppointmentResponse, error)
	Update(ctx context.Context, appointment domain.AppointmentRequest, id int) (domain.AppointmentRequest, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, appointment domain.AppointmentRequest, id int) (domain.AppointmentRequest, error)
}

type service struct {
	repository RepositoryAppointments
}

func NewServiceAppointments(repository RepositoryAppointments) ServiceAppointments {
	return &service{repository: repository}
}

// Create is a method that creates a new appointment.
func (s *service) Create(ctx context.Context, appointment domain.AppointmentRequest) (domain.AppointmentRequest, error) {
	appointment, err := s.repository.Create(ctx, appointment)
	if err != nil {
		log.Println("[AppointmentsService][Create] error creating appointment", err)
		return domain.AppointmentRequest{}, err
	} 
    return appointment, nil;
}

// GetAll is a method that returns all appointment.
func (s *service) GetAll(ctx context.Context) ([]domain.AppointmentResponse, error) {
	appointments, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[AppointmentsService][GetAll] error getting all appointments", err)
		return []domain.AppointmentResponse{}, err
	}

	return appointments, nil
}

// GetByID is a method that returns a appointment by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.AppointmentResponse, error) {
	appointment, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[AppointmentsService][GetByID] error getting appointment by ID", err)
		return domain.AppointmentResponse{}, err
	}

	return appointment, nil
}

// Update is a method that updates a appointment by ID.
func (s *service) Update(ctx context.Context, appointment domain.AppointmentRequest, id int) (domain.AppointmentRequest, error) {
	appointment, err := s.repository.Update(ctx, appointment, id)
	if err != nil {
		log.Println("[AppointmentsService][Update] error updating appointment by ID", err)
		return domain.AppointmentRequest{}, err
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
func (s *service) Patch(ctx context.Context, appointment domain.AppointmentRequest, id int) (domain.AppointmentRequest, error) {
	// ALGO ASI
	// appointmentStore, err := s.repository.GetByID(ctx, id)
	// if err != nil {
	// 	log.Println("[AppointmentService][Patch] error getting appointment by ID", err)
	// 	return domain.AppointmentRequest{}, err
	// }

	// appointmentPatch, err := s.validatePatch(appointmentStore, appointment)
	// if err != nil {
	// 	log.Println("[AppointmentService][Patch] error validating appointment", err)
	// 	return domain.AppointmentRequest{}, err
	// }
	
	appointment, err := s.repository.Patch(ctx, appointment, id)
	if err != nil {
		log.Println("[AppointmentsService][Patch] error patching appointment by ID", err)
		return domain.AppointmentRequest{}, err
	}

	return appointment, nil
}

// hay que validar el patch utilizando esta funcion que devuelve un response que tiene turnoid, dentistid, patientid, 
// fecha, hora, descripcion y que compara con el request que se manda por parametro en la funcion patch y no darle bola 
//a dentistlastname y patientlastname IR A LINEA 82
func (s *service) validatePatch(appointmentStore domain.AppointmentResponse, appointment domain.AppointmentResponse) (domain.AppointmentResponse, error) {
	if appointmentStore.Id != appointment.Id {
		return domain.AppointmentResponse{}, domain.ErrInvalidID
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