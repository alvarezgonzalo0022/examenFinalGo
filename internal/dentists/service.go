package dentists

import (
	"context"
	"log"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type ServiceDentists interface {
	Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error)
	GetAll(ctx context.Context) ([]domain.Dentist, error)
	GetByID(ctx context.Context, id int) (domain.Dentist, error)
	Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error)
}

type service struct {
	repository RepositoryDentists
}

func NewServiceDentist(repository RepositoryDentists) ServiceDentists {
	return &service{repository: repository}
}

// Create is a method that creates a new odontologo.
func (s *service) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.repository.Create(ctx, dentist)
	if err != nil {
		log.Println("[OdontologosService][Create] error creating dentist", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// GetAll is a method that returns all odontologos.
func (s *service) GetAll(ctx context.Context) ([]domain.Dentist, error) {
	listDentists, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[OdontologosService][GetAll] error getting all odontologos", err)
		return []domain.Dentist{}, err
	}

	return listDentists, nil
}

// GetByID is a method that returns a odontologo by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	dentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[OdontologosService][GetByID] error getting odontologo by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update is a method that updates a odontologo by ID.
func (s *service) Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	dentist, err := s.repository.Update(ctx, dentist, id)
	if err != nil {
		log.Println("[OdontologosService][Update] error updating odontologo by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Delete is a method that deletes a odontologo by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[OdontologosService][Delete] error deleting odontologo by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a odontologo by ID.
func (s *service) Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	odontologoStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[OdontologosService][Patch] error getting odontologo by ID", err)
		return domain.Dentist{}, err
	}

	odontologoPatch, err := s.validatePatch(odontologoStore, dentist)
	if err != nil {
		log.Println("[OdontologosService][Patch] error validating odontologo", err)
		return domain.Dentist{}, err
	}

	dentist, err = s.repository.Patch(ctx, odontologoPatch, id)
	if err != nil {
		log.Println("[OdontologosService][Patch] error patching odontologo by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(odontologoStore, dentist domain.Dentist) (domain.Dentist, error) {

	if dentist.FirstName != "" {
		odontologoStore.FirstName = dentist.FirstName
	}

	if dentist.LastName != "" {
		odontologoStore.LastName = dentist.LastName
	}

	if dentist.RegistrationId != "" {
		odontologoStore.RegistrationId = dentist.RegistrationId
	}

	return odontologoStore, nil

}
