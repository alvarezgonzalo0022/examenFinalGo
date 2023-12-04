package dentists

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
	ErrNotFound = errors.New("dentist not found")
)

type repositorydentistsmysql struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepository(db *sql.DB) RepositoryDentists {
	return &repositorydentistsmysql{db: db}
}

// Create is a method that creates a new dentist.
func (r *repositorydentistsmysql) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryInsertDentist)
	if err != nil {
		return domain.Dentist{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.FirstName,
		dentist.LastName,
		dentist.RegistrationId,
	)

	if err != nil {
		return domain.Dentist{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Dentist{}, ErrLastInsertedId
	}

	dentist.Id = int(lastId)

	return dentist, nil

}

// GetAll is a method that returns all dentists.
func (r *repositorydentistsmysql) GetAll(ctx context.Context) ([]domain.Dentist, error) {
	rows, err := r.db.Query(QueryGetAllDentists)
	if err != nil {
		return []domain.Dentist{}, err
	}

	defer rows.Close()

	var dentists []domain.Dentist

	for rows.Next() {
		var dentist domain.Dentist
		err := rows.Scan(
			&dentist.Id,
			&dentist.FirstName,
			&dentist.LastName,
			&dentist.RegistrationId,
		)
		if err != nil {
			return []domain.Dentist{}, err
		}

		dentists = append(dentists, dentist)
	}

	if err := rows.Err(); err != nil {
		return []domain.Dentist{}, err
	}

	return dentists, nil
}

// GetByID is a method that returns a dentist by ID.
func (r *repositorydentistsmysql) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	row := r.db.QueryRow(QueryGetDentistById, id)

	var dentist domain.Dentist
	err := row.Scan(
		&dentist.Id,
		&dentist.FirstName,
		&dentist.LastName,
		&dentist.RegistrationId,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update is a method that updates a dentist by ID.
func (r *repositorydentistsmysql) Update(
	ctx context.Context,
	dentist domain.Dentist,
	id int) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.FirstName,
		dentist.LastName,
		dentist.RegistrationId,
		id,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	if rowsAffected == 0 {
		return domain.Dentist{}, fmt.Errorf("no se encontró ningún dentista con ID %d", id)
	}

	dentist.Id = id

	return dentist, nil

}

// Delete is a method that deletes a dentist by ID.
func (r *repositorydentistsmysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteDentist, id)
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

// Patch is a method that updates a dentist by ID.
func (r *repositorydentistsmysql) Patch(
	ctx context.Context,
	dentist domain.Dentist,
	id int) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.FirstName,
		dentist.LastName,
		dentist.RegistrationId,
		id,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	if rowsAffected == 0 {
		return domain.Dentist{}, fmt.Errorf("no se encontró ningún dentista con ID %d", id)
	}

	return dentist, nil
}
