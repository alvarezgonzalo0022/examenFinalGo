package turnos

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
	"github.com/alvarezgonzalo0022/examenFinalGo/internal/turnos"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/web"
	"github.com/gin-gonic/gin"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("paciente not found")
)


type repositoryturnomysql struct {
	db *sql.DB;
}

func NewMySqlRepository(db *sql.DB) RepositoryTurnos {
	return &repositoryturnomysql{db: db}
}

func (r *repositoryturnomysql) Create(ctx *gin.Context, turno domain.Turno) (domain.Turno, error) {
	statement, err := r.db.Prepare(QueryInsertTurno)
	if err != nil {
		return domain.Turno{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		turno.IdDentist,
		turno.IdPatient,
		turno.Date,
		turno.Time,
	)

	if err != nil {
		return domain.Turno{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Turno{}, ErrLastInsertedId
	}

	turno.Id = int(lastId)

	return turno, nil
}

func (r *repositoryturnomysql) GetAll(ctx *gin.Context) ([]domain.Turno, error) {
	rows, err := r.db.Query(QueryGetAllTurnos)
	if err != nil {
		return nil, ErrQuery
	}

	defer rows.Close()

	var turnos []domain.Turno

	for rows.Next() {
		var turno domain.Turno

		err := rows.Scan(
			&turno.Id,
			&turno.IdDentist,
			&turno.IdPatient,
			&turno.Date,
			&turno.Time,
		)

		if err != nil {
			log.Println(err)
			continue
		}

		turnos = append(turnos, turno)
	}

	return turnos, nil
}


//
func (r *repositoryturnomysql) GetById(ctx *gin.Context, id int) (domain.Turno, error) {
	statement, err := r.db.Prepare(QueryGetTurnoById)
	if err != nil {
		return domain.Turno{}, ErrPrepareStatement
	}

	defer statement.Close()

	row := statement.QueryRow(id)

	var turno domain.Turno

	err = row.Scan(
		&turno.Id,
		&turno.IdDentist,
		&turno.IdPatient,
		&turno.AppointmentDate,
		&turno.AppointmentTime,
		&turno.Description,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Turno{}, domain.ErrAppointmentNotFound
		}
		return domain.Turno{}, domain.ErrScanRow
	}

	return turno, nil
}



func (r *repositoryturnomysql) Update(ctx context.Context, turno domain.Turno, id int) (domain.Turno, error) {
    statement, err := r.db.Prepare(QueryUpdateTurno)
    if err != nil {
        return domain.Turno{}, err
    }

    defer statement.Close()

    result, err := statement.Exec(
		&turno.IdDentist,
		&turno.IdPatient,
		&turno.AppointmentDate,
		&turno.AppointmentTime,
		&turno.Description,
    )

    if err != nil {
        return domain.Turno{}, err
    }

    _, err = result.RowsAffected()
    if err != nil {
        return domain.Turno{}, err
    }

    turno.Id = id

    return turno, nil
}


func (r *repositoryturnomysql) Delete(ctx context.Context, id int) error {
    result, err := r.db.Exec(QueryDeleteTurno, id)
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
