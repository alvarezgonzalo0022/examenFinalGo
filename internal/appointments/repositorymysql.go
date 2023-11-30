package appointments

import (
	"context"
	"database/sql"
	"errors"
	"time"

	// "time"

	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	// "github.com/alvarezgonzalo0022/examenFinalGo/internal/appointments"
	// "github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrEmpty            = errors.New("empty list")
	ErrNotFound         = errors.New("paciente not found")
)

type repositoryappointmentmysql struct {
	db *sql.DB
}

func NewMySqlRepository(db *sql.DB) RepositoryAppointments {
	return &repositoryappointmentmysql{db: db}
}

func (r *repositoryappointmentmysql) Create(ctx context.Context, appointment domain.AppointmentRequest) (domain.AppointmentRequest, error) {
	statement, err := r.db.Prepare(QueryInsertAppointment)
	if err != nil {
		return domain.AppointmentRequest{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		appointment.IdDentist,
		appointment.IdPatient,
		appointment.AppointmentDate,
		appointment.AppointmentTime,
		appointment.Description,
	)

	if err != nil {
		return domain.AppointmentRequest{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.AppointmentRequest{}, ErrLastInsertedId
	}

	appointment.Id = int(lastId)

	return appointment, nil
}

func (r *repositoryappointmentmysql) GetAll(ctx context.Context) ([]domain.AppointmentResponse, error) {
   // query := `SELECT appointments.*, dentists.last_name AS dentist_lastname, patients.last_name AS patient_lastname FROM appointments INNER JOIN dentists ON appointments.dentist_id = dentists.id INNER JOIN patients ON appointments.patient_id = patients.id `


	rows, err := r.db.Query(QueryGetAllAppointment)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var appointments []domain.AppointmentResponse

    for rows.Next() {
        var appointment domain.AppointmentResponse
        var appointmentTimeStr string  // Variable para almacenar la cadena de tiempo

        err := rows.Scan(&appointment.Id, &appointment.IdDentist, &appointment.IdPatient,
            &appointment.AppointmentDate, &appointmentTimeStr, &appointment.Description, &appointment.DentistLastName, &appointment.PatientLastName)
        if err != nil {
            return nil, err
        }

        // Convierte la cadena de tiempo a un objeto time.Time
        appointment.AppointmentTime, err = time.Parse("15:04:05", appointmentTimeStr)
        if err != nil {
            return nil, err
        }

        appointments = append(appointments, appointment)
    }

    return appointments, nil

}

func (r *repositoryappointmentmysql) GetByID(ctx context.Context, id int) (domain.AppointmentResponse, error) {
    row := r.db.QueryRow(QueryGetByIdAppointment, id)

    var appointment domain.AppointmentResponse
    var appointmentTimeStr string 

    err := row.Scan(
        &appointment.Id,
        &appointment.IdDentist,
        &appointment.IdPatient,
        &appointment.AppointmentDate,
        &appointmentTimeStr, 
        &appointment.Description,
		&appointment.PatientLastName,
		&appointment.DentistLastName,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return domain.AppointmentResponse{}, ErrNotFound
        }
        return domain.AppointmentResponse{}, err
    }


    if appointmentTimeStr != "" {
        parsedTime, err := time.Parse("15:04:05", appointmentTimeStr) 
        if err != nil {
            return domain.AppointmentResponse{}, err
        }
        appointment.AppointmentTime = parsedTime
    }

    return appointment, nil
}




func (r *repositoryappointmentmysql) Update(ctx context.Context, appointment domain.AppointmentRequest, id int) (domain.AppointmentRequest, error) {
	statement, err := r.db.Prepare(QueryUpdateAppointment)
	if err != nil {
		return domain.AppointmentRequest{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		appointment.IdDentist,
		appointment.IdPatient,
		appointment.AppointmentDate,
		appointment.AppointmentTime,
		appointment.Description,
	)

	if err != nil {
		return domain.AppointmentRequest{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.AppointmentRequest{}, err
	}

	appointment.Id = id

	return appointment, nil
}

func (r *repositoryappointmentmysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteAppointment, id)
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

func (r *repositoryappointmentmysql) Patch(
	ctx context.Context,
	appointment domain.AppointmentRequest,
	id int) (domain.AppointmentRequest, error) {
	statement, err := r.db.Prepare(QueryUpdateAppointment)
	if err != nil {
		return domain.AppointmentRequest{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		appointment.IdDentist,
		appointment.IdPatient,
		appointment.AppointmentDate,
		appointment.AppointmentTime,
		appointment.Description,
	)

	if err != nil {
		return domain.AppointmentRequest{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.AppointmentRequest{}, err
	}

	return appointment, nil
}
