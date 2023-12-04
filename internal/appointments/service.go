package appointments

import (
	"context"
	"log"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/dentists"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/patients"
)

type ServiceAppointments interface {
	Create(ctx context.Context, request domain.AppointmentRequest) (domain.AppointmentRequest, error)
	GetAll(ctx context.Context) ([]domain.AppointmentResponse, error)
	GetByID(ctx context.Context, id int) (domain.AppointmentResponse, error)
	Update(ctx context.Context, appointment domain.AppointmentRequest, id int) (domain.AppointmentRequest, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, appointment domain.AppointmentPatchRequest, id int) (domain.AppointmentPatchRequest, error)
}

type service struct {
	repository RepositoryAppointments
	dentistService dentists.ServiceDentists
	patientService patients.ServicePatients
}

func NewServiceAppointments(repository RepositoryAppointments,
	dentistService dentists.ServiceDentists,
	patientService patients.ServicePatients,) ServiceAppointments {
	return &service{
		repository: repository,
		dentistService:  dentistService,
		patientService:  patientService,
	}
}

// Create is a method that creates a new appointment.
func (s *service) Create(ctx context.Context, appointment domain.AppointmentRequest) (domain.AppointmentRequest, error) {
	// Validar la existencia del dentista
	_, err := s.dentistService.GetByID(ctx, appointment.DentistId)
	if err != nil {
		log.Println("[AppointmentsService][Create] error validating dentist existence", err)
		return domain.AppointmentRequest{}, err
	}
	// Validar la existencia del paciente
	_, err = s.patientService.GetByID(ctx, appointment.PatientId)
	if err != nil {
		log.Println("[AppointmentsService][Create] error validating patient existence", err)
		return domain.AppointmentRequest{}, err
	}
	appointmentResult, err := s.repository.Create(ctx, appointment)
	if err != nil {
		log.Println("[AppointmentsService][Create] error creating appointment", err)
		return domain.AppointmentRequest{}, err
	} 
    return appointmentResult, nil;
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
	_, err := s.dentistService.GetByID(ctx, appointment.DentistId)
	if err != nil {
		log.Println("[AppointmentsService][Create] error validating dentist existence", err)
		return domain.AppointmentRequest{}, err
	}
	// Validar la existencia del paciente
	_, err = s.patientService.GetByID(ctx, appointment.PatientId)
	if err != nil {
		log.Println("[AppointmentsService][Create] error validating patient existence", err)
		return domain.AppointmentRequest{}, err
	}
	appointment, err = s.repository.Update(ctx, appointment, id)
	if err != nil {
		log.Println("[AppointmentsService][Update] error updating appointment by ID", err)
		return domain.AppointmentRequest{}, err
	}

	return appointment, nil
}

// Delete is a method that deletes a appointment by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	_, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[AppointmentsService][GetByID] error validating appointment existence", err)
		return err
	}

	err = s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[AppointmentsService][Delete] error deleting appointment by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates an appointment by ID.
func (s *service) Patch(ctx context.Context, appointment domain.AppointmentPatchRequest, id int) (domain.AppointmentPatchRequest, error) {
    // Validar la existencia del dentista
	if appointment.DentistId != 0 {
		_, err := s.dentistService.GetByID(ctx, appointment.DentistId)
		if err != nil {
			log.Println("[Service][Patch] error validating dentist existence", err)
			return domain.AppointmentPatchRequest{}, err
		}
    }

    // Validar la existencia del paciente
	if appointment.PatientId != 0 {
		_, err := s.patientService.GetByID(ctx, appointment.PatientId)
		if err != nil {
			log.Println("[Service][Patch] error validating patient existence", err)
			return domain.AppointmentPatchRequest{}, err
		}
    }

    appointmentStore, err := s.repository.GetByID(ctx, id)
    if err != nil {
        log.Println("[AppointmentsService][Patch] error getting appointment by ID", err)
        return domain.AppointmentPatchRequest{}, err
    }

    // Actualizar solo los campos proporcionados en la solicitud
    appointment, err = s.validatePatch(appointmentStore, appointment)
    if err != nil {
        log.Println("[AppointmentsService][Patch] error validating appointment", err)
		log.Println(appointment)
        return domain.AppointmentPatchRequest{}, err
    }

    appointment, err = s.repository.Patch(ctx, appointment, id)
    if err != nil {
        log.Println("[AppointmentsService][Patch] error patching appointment by ID", err)
        return domain.AppointmentPatchRequest{}, err
    }

    return appointment, nil
}

// validatePatch actualiza los campos en appointment con los valores de appointmentStore si están presentes en la solicitud.
func (s *service) validatePatch(appointmentStore domain.AppointmentResponse, appointment domain.AppointmentPatchRequest) (domain.AppointmentPatchRequest, error) {
    // Comparar y actualizar cada campo si está presente en la solicitud
    if appointment.AppointmentDate == "" && appointmentStore.AppointmentDate != "" {
        appointment.AppointmentDate = appointmentStore.AppointmentDate
    }

    if appointment.AppointmentTime == "" && appointmentStore.AppointmentTime != "" {
        appointment.AppointmentTime = appointmentStore.AppointmentTime
    }

    if appointment.PatientId == 0 && appointmentStore.PatientId != 0 {
        appointment.PatientId = appointmentStore.PatientId
    }

    if appointment.DentistId == 0 && appointmentStore.DentistId != 0 {
        appointment.DentistId = appointmentStore.DentistId
    }

    if appointment.Description == "" && appointmentStore.Description != "" {
        appointment.Description = appointmentStore.Description
    }

    // Devolver la versión actualizada de appointment
    return appointment, nil
}


