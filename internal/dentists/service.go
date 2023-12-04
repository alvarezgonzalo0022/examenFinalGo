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

// Create is a method that creates a new dentist.
func (s *service) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.repository.Create(ctx, dentist)
	if err != nil {
		log.Println("[DentistService][Create] error creating dentist", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// GetAll is a method that returns all dentists.
func (s *service) GetAll(ctx context.Context) ([]domain.Dentist, error) {
	listDentists, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[DentistsService][GetAll] error getting all dentists", err)
		return []domain.Dentist{}, err
	}

	return listDentists, nil
}

// GetByID is a method that returns a dentist by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	dentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistsService][GetByID] error getting dentist by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update is a method that updates a dentist by ID.
func (s *service) Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	dentist, err := s.repository.Update(ctx, dentist, id)
	if err != nil {
		log.Println("[DentistsService][Update] error updating dentist by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Delete is a method that deletes a dentist by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[DentistsService][Delete] error deleting dentist by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a dentist by ID.
func (s *service) Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	dentistStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistsService][Patch] error getting dentist by ID", err)
		return domain.Dentist{}, err
	}

	dentistPatch, err := s.validatePatch(dentistStore, dentist)
	if err != nil {
		log.Println("[DentistsService][Patch] error validating dentist", err)
		return domain.Dentist{}, err
	}

	dentist, err = s.repository.Patch(ctx, dentistPatch, id)
	if err != nil {
		log.Println("[DentistsService][Patch] error patching dentist by ID", err)
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(dentistStore, dentist domain.Dentist) (domain.Dentist, error) {

	if dentist.FirstName != "" {
		dentistStore.FirstName = dentist.FirstName
	}

	if dentist.LastName != "" {
		dentistStore.LastName = dentist.LastName
	}

	if dentist.RegistrationId != "" {
		dentistStore.RegistrationId = dentist.RegistrationId
	}

	return dentistStore, nil

}
