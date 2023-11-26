package products

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("product not found")
)

type repositorymysql struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepository(db *sql.DB) Repository {
	return &repositorymysql{db: db}
}

// Create is a method that creates a new odontologo.
func (r *repositorymysql) Create(ctx context.Context, odontologo domain.Odontologo) (domain.Odontologo, error) {
	statement, err := r.db.Prepare(QueryInsertOdontologo)
	if err != nil {
		return domain.Odontologo{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Name,
		odontologo.LastName,
		odontologo.Code,
	)

	if err != nil {
		return domain.Odontologo{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Odontologo{}, ErrLastInsertedId
	}

	odontologo.Id = int(lastId)

	return odontologo, nil

}

// GetAll is a method that returns all odontologos.
func (r *repositorymysql) GetAll(ctx context.Context) ([]domain.Odontologo, error) {
	rows, err := r.db.Query(QueryGetAllOdontologos)
	if err != nil {
		return []domain.Odontologo{}, err
	}

	defer rows.Close()

	var odontologos []domain.Odontologo

	for rows.Next() {
		var odontologo domain.Odontologo
		err := rows.Scan(
			&odontologo.Id,
			&odontologo.Name,
			&odontologo.LastName,
			&odontologo.Code,
		)
		if err != nil {
			return []domain.Odontologo{}, err
		}

		odontologos = append(odontologos, odontologo)
	}

	if err := rows.Err(); err != nil {
		return []domain.Odontologo{}, err
	}

	return odontologos, nil
}

// GetByID is a method that returns a odontologo by ID.
func (r *repositorymysql) GetByID(ctx context.Context, id int) (domain.Odontologo, error) {
	row := r.db.QueryRow(QueryGetOdontologoById, id)

	var odontologo domain.Odontologo
	err := row.Scan(
		&odontologo.Id,
		&odontologo.Name,
		&odontologo.LastName,
		&odontologo.Code,
	)

	if err != nil {
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}

// Update is a method that updates a odontologo by ID.
func (r *repositorymysql) Update(
	ctx context.Context,
	odontologo domain.Odontologo,
	id int) (domain.Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateOdontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Name,
		odontologo.LastName,
		odontologo.Code,
	)

	if err != nil {
		return domain.Odontologo{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Odontologo{}, err
	}

	odontologo.Id = id

	return odontologo, nil

}

// Delete is a method that deletes a product by ID.
func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteOdontologo, id)
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

// Patch is a method that updates a odontologo by ID.
func (r *repositorymysql) Patch(
	ctx context.Context,
	odontologo domain.Odontologo,
	id int) (domain.Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateOdontologo)
	if err != nil {
		return domain.Odontologo{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		odontologo.Name,
		odontologo.LastName,
		odontologo.Code,
	)

	if err != nil {
		return domain.Odontologo{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Odontologo{}, err
	}

	return odontologo, nil
}
