package domain

import (
	"time"

	"github.com/alvarezgonzalo0022/examenFinalGo/internal/domain"
)

type Appointment struct {
	Id              int   		    `json:"id"`
	Dentist      	domain.Dentist  `json:"id_dentist"  binding:"required"`
	Patient     	domain.Patient  `json:"id_patient"  binding:"required"`
	AppointmentDate time.Time 		`json:"appointment_date"`
	AppointmentTime time.Time 		`json:"appointment_time"`
	Description     string			`json:"description" binding:"required"`
}

type AppointmentRequest struct {
	Id              int       `json:"id"`
	IdDentist       int       `json:"id_dentist"  binding:"required"`
	IdPatient       int       `json:"id_patient"  binding:"required"`
	AppointmentDate time.Time `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	Description     string    `json:"description" binding:"required"`
}

type AppointmentResponse struct {
	Id              int       `json:"id"`
	IdDentist       int       `json:"id_dentist"  binding:"required"`
	IdPatient       int       `json:"id_patient"  binding:"required"`
	AppointmentDate time.Time `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	Description     string    `json:"description" binding:"required"`
	PatientLastName string    `json:"patient_lastname"`
	DentistLastName string    `json:"dentist_lastname"`
}
