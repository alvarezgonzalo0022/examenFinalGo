package appointments

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/alvarezgonzalo0022/examenFinalGo/internal/dentists"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/patients"
	// "github.com/alvarezgonzalo0022/examenFinalGo/internal/appointments"
	// "github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrEmpty            = errors.New("empty list")
	ErrNotFound         = errors.New("appointment not found")
)

type repositoryappointmentmysql struct {
	db *sql.DB
	dentistService	dentists.ServiceDentists
	patientService	patients.ServicePatients
}

func NewMySqlRepository(db *sql.DB, dentistService dentists.ServiceDentists,
	patientService	patients.ServicePatients) RepositoryAppointments {
	return &repositoryappointmentmysql{
		db: 		db,
		dentistService: dentistService,
		patientService: patientService,
	}
}

func (r *repositoryappointmentmysql) Create(ctx context.Context, appointment domain.AppointmentRequest) (domain.AppointmentRequest, error) {
	statement, err := r.db.Prepare(QueryInsertAppointment)
	if err != nil {
		return domain.AppointmentRequest{}, ErrPrepareStatement
	}

	defer statement.Close()

	appointmentDate, err := time.Parse("2006-01-02", appointment.AppointmentDate)
    if err != nil {
        log.Println("Error parsing appointment date:", err)
        return domain.AppointmentRequest{}, domain.ErrInvalidDateFormat
    }
	appointmentDate = time.Date(appointmentDate.Year(), appointmentDate.Month(), appointmentDate.Day(), 0, 0, 0, 0, time.UTC)

	result, err := statement.Exec(
		appointment.PatientId,
		appointment.DentistId,
		appointmentDate,
		appointment.AppointmentTime,
		appointment.Description,
	)

	if err != nil {
		log.Println("Error:", err)
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
    rows, err := r.db.Query(QueryGetAllAppointment)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var appointments []domain.AppointmentResponse

    for rows.Next() {
        var appointment domain.AppointmentResponse
        var appointmentDateStr, appointmentTimeStr string
        err := rows.Scan(
			&appointment.Id, 
			&appointment.DentistId,
			&appointment.PatientId,
            &appointmentDateStr,
			&appointmentTimeStr,
			&appointment.Description,
			&appointment.DentistLastName,
			&appointment.PatientLastName)
        if err != nil {
            return nil, err
        }

        appointment.AppointmentDate = appointmentDateStr
        appointment.AppointmentTime = appointmentTimeStr

        appointments = append(appointments, appointment)
    }

    return appointments, nil
}

func (r *repositoryappointmentmysql) GetByID(ctx context.Context, id int) (domain.AppointmentResponse, error) {
    row := r.db.QueryRow(QueryGetByIdAppointment, id)

	
	
	var appointment domain.AppointmentResponse
    err := row.Scan(
        &appointment.Id,
        &appointment.DentistId,
        &appointment.PatientId,
        &appointment.AppointmentDate,
        &appointment.AppointmentTime,
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

	// appointmentDate, err := time.Parse("2006-01-02", appointment.AppointmentDate)
    // if err != nil {
    //     log.Println("Error parsing appointment date:", err)
    //     return domain.AppointmentResponse{}, domain.ErrInvalidDateFormat
    // }

	appointment.AppointmentDate = strings.Split(appointment.AppointmentDate, "T")[0]

    return appointment, nil
}


func (r *repositoryappointmentmysql) Update(ctx context.Context, appointment domain.AppointmentRequest, id int) (domain.AppointmentRequest, error) {
	statement, err := r.db.Prepare(QueryUpdateAppointment)
	if err != nil {
		return domain.AppointmentRequest{}, err
	}

	defer statement.Close()

	// appointmentDate, err := time.Parse("2006-01-02", appointment.AppointmentDate)
    // if err != nil {
    //     log.Println("Error parsing appointment date:", err)
    //     return domain.AppointmentRequest{}, domain.ErrInvalidDateFormat
    // }
	// appointmentDate = time.Date(appointmentDate.Year(), appointmentDate.Month(), appointmentDate.Day(), 0, 0, 0, 0, time.UTC)

	appointment.AppointmentDate = strings.Split(appointment.AppointmentDate, "T")[0]

	appointmentTime, err := time.Parse("15:04", appointment.AppointmentTime)
    if err != nil {
        // Manejar el error según tus necesidades
        log.Println("Error parsing appointment time:", err)
        return domain.AppointmentRequest{}, domain.ErrInvalidTimeFormat
    }
	appointmentTime = time.Date(1, 1, 1, appointmentTime.Hour(), appointmentTime.Minute(), 0, 0, time.UTC)

	result, err := statement.Exec(
		appointment.DentistId,
		appointment.PatientId,
		appointment.AppointmentDate,
		appointmentTime,
		appointment.Description,
		id,
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
	appointment domain.AppointmentPatchRequest,
	id int) (domain.AppointmentPatchRequest, error) {
	statement, err := r.db.Prepare(QueryUpdateAppointment)
	if err != nil {
		return domain.AppointmentPatchRequest{}, err
	}
	defer statement.Close()

	// var appointmentDate time.Time
	if appointment.AppointmentDate != "" {
	// appointmentDate, err := time.Parse("2006-01-02", appointment.AppointmentDate)
	// if err != nil {
	// 	log.Println("Error parsing appointment date:", err)
	// 	return domain.AppointmentPatchRequest{}, domain.ErrInvalidDateFormat
	// 	}
		// appointmentDate = time.Date(appointmentDate.Year(), appointmentDate.Month(), appointmentDate.Day(), 0, 0, 0, 0, time.UTC)
		appointment.AppointmentDate = strings.Split(appointment.AppointmentDate, "T")[0]
	}
	



	result, err := statement.Exec(
		appointment.DentistId,
		appointment.PatientId,
		appointment.AppointmentDate,
		appointment.AppointmentTime,
		appointment.Description,
		id,
	)

	if err != nil {
		log.Println("Error updating appointment:", err)
		return domain.AppointmentPatchRequest{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return domain.AppointmentPatchRequest{}, err
	}

	if rowsAffected == 0 {
		return domain.AppointmentPatchRequest{}, sql.ErrNoRows
	}

	appointment.Id = id

	return appointment, nil
}
