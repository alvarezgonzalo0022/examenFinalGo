package patients

import (
	"context"
	"log"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type ServicePatients interface {
	Create(ctx context.Context, patient domain.Patient) (domain.Patient, error)
	GetAll(ctx context.Context) ([]domain.Patient, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
}

type service struct {
	repository RepositoryPatients
}

func NewServicePatients(repository RepositoryPatients) ServicePatients {
	return &service{repository: repository}
}

// Create is a method that creates a new patient.
func (s *service) Create(ctx context.Context, patient domain.Patient) (domain.Patient, error) {
	patient, err := s.repository.Create(ctx, patient)
	if err != nil {
		log.Println("[PatientsService][Create] error creating patient", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// GetAll is a method that returns all patients.
func (s *service) GetAll(ctx context.Context) ([]domain.Patient, error) {
	patients, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[PatientsService][GetAll] error getting all patients", err)
		return []domain.Patient{}, err
	}

	return patients, nil
}

// GetByID is a method that returns a patient by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	patient, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PatientsService][GetByID] error getting patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Update is a method that updates a patient by ID.
func (s *service) Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	patient, err := s.repository.Update(ctx, patient, id)
	if err != nil {
		log.Println("[PatientsService][Update] error updating patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Delete is a method that deletes a patient by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[PatientsService][Delete] error deleting patient by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a patient by ID.
func (s *service) Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	patientStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PatientsService][Patch] error getting patient by ID", err)
		return domain.Patient{}, err
	}

	patientPatch, err := s.validatePatch(patientStore, patient)
	if err != nil {
		log.Println("[PatientsService][Patch] error validating patient", err)
		return domain.Patient{}, err
	}

	patient, err = s.repository.Patch(ctx, patientPatch, id)
	if err != nil {
		log.Println("[PatientsService][Patch] error patching patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(patientStore, patient domain.Patient) (domain.Patient, error) {

	if patient.FirstName != "" {
		patientStore.FirstName = patient.FirstName
	}

	if patient.LastName != "" {
		patientStore.LastName = patient.LastName
	}

	if patient.Address != "" {
		patientStore.Address = patient.Address
	}

	if patient.Document != "" {
		patientStore.Document = patient.Document
	}

	if !patient.AdmissionDate.IsZero() {
		patientStore.AdmissionDate = patient.AdmissionDate
	}

	return patientStore, nil
}
