package domain

import "time"

type Turno struct {
	Id           int       `json:"id"`
	IdOdontologo int       `json:"id_odontologo"  binding:"required"`
	IdPaciente   int       `json:"id_paciente"  binding:"required"`
	Description  string    `json:"description" binding:"required"`
	FechaTurno   time.Time `json:"fecha_turno"`
	HoraTurno    string    `json:"hora_turno"`
}
