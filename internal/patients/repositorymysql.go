package patients

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("patient not found")
)

type repositorypatientsmysql struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewPatientsMySqlRepository(db *sql.DB) RepositoryPatients {
	return &repositorypatientsmysql{db: db}
}

// Create is a method that creates a new patient.
func (r *repositorypatientsmysql) Create(ctx context.Context, patient domain.Patient) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryInsertPatient)
	if err != nil {
		return domain.Patient{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.FirstName,
		patient.LastName,
		patient.Address,
		patient.Document,
		patient.AdmissionDate,
	)

	if err != nil {
		return domain.Patient{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Patient{}, ErrLastInsertedId
	}

	patient.Id = int(lastId)

	return patient, nil

}

// GetAll is a method that returns all patients.
func (r *repositorypatientsmysql) GetAll(ctx context.Context) ([]domain.Patient, error) {
	rows, err := r.db.Query(QueryGetAllPatients)
	if err != nil {
		return []domain.Patient{}, err
	}

	defer rows.Close()

	var patients []domain.Patient

	for rows.Next() {
		var patient domain.Patient
		err := rows.Scan(
			&patient.Id,
			&patient.FirstName,
			&patient.LastName,
			&patient.Address,
			&patient.Document,
			&patient.AdmissionDate,
		)
		if err != nil {
			return []domain.Patient{}, err
		}

		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return []domain.Patient{}, err
	}

	return patients, nil
}

// GetByID is a method that returns a patient by ID.
func (r *repositorypatientsmysql) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	row := r.db.QueryRow(QueryGetPatientById, id)
	if err := row.Err(); err != nil {
		return domain.Patient{}, err
	}

	var patient domain.Patient

	err := row.Scan(
		&patient.Id,
		&patient.FirstName,
		&patient.LastName,
		&patient.Address,
		&patient.Document,
		&patient.AdmissionDate,
	)

	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            // Devolver un error específico de tu dominio indicando que el paciente no se encontró
            return domain.Patient{}, domain.ErrPatientNotFound
        }
        // Otros errores relacionados con la base de datos
        return domain.Patient{}, fmt.Errorf("error scanning row %v", err)
    }

	return patient, nil
}

// Update is a method that updates a patient by ID.
func (r *repositorypatientsmysql) Update(
	ctx context.Context,
	patient domain.Patient,
	id int) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryUpdatePatient)
	if err != nil {
		return domain.Patient{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.FirstName,
		patient.LastName,
		patient.Address,
		patient.Document,
		patient.AdmissionDate,
		id,
	)

	if err != nil {
		return domain.Patient{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	if rowsAffected == 0 {
		return domain.Patient{}, fmt.Errorf("no se encontró ningún paciente con ID %d", id)
	}

	patient.Id = id

	return patient, nil

}

// Delete is a method that deletes a patient by ID.
func (r *repositorypatientsmysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeletePatient, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}

// Patch is a method that updates a patient by ID.
func (r *repositorypatientsmysql) Patch(
	ctx context.Context,
	patient domain.Patient,
	id int) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryUpdatePatient)
	if err != nil {
		return domain.Patient{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.FirstName,
		patient.LastName,
		patient.Address,
		patient.Document,
		patient.AdmissionDate,
	)

	if err != nil {
		return domain.Patient{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	if rowsAffected == 0 {
		return domain.Patient{}, fmt.Errorf("no se encontró ningún paciente con ID %d", id)
	}

	return patient, nil
}
