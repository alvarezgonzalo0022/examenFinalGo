package domain

import "time"

type Turno struct {
	Id              int    `json:"id"`
	Description     string `json:"description" binding:"required"`
	DentistLicense  string `json:"dentist_license" binding:"required"`
	PatientIdentity string `json:"patient_identity" binding:"required"`

	IdOdontologo int       `json:"id_odontologo"`
	IdPaciente   int       `json:"id_paciente"`
	Date         time.Time `json:"date"`
	Time         string    `json:"time"`
}
