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
		turno.IdOdontologo,
		turno.IdPaciente,
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
			&turno.IdOdontologo,
			&turno.IdPaciente,
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
		&turno.IdOdontologo,
		&turno.IdPaciente,
		&turno.Date,
		&turno.Time,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Turno{}, ErrTurnoNotFound
		}
		return domain.Turno{}, ErrScanRow
	}

	return turno, nil
}
	if err != nil {
		return domain.Turno{}, ErrPrepareStatement
	}

	defer statement.Close()

	row := statement.QueryRow(id)

	var turno domain.Turno

	err = row.Scan(
		&turno.Id,
		&turno.IdOdontologo,
		&turno.IdPaciente,
		&turno.Date,
		&turno.Time,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Turno{}, ErrTurnoNotFound
		}
		return domain.Turno{}, ErrScanRow
	}

	return turno, nil
}
	