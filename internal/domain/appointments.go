package domain

import "time"

type Appointment struct {
	Id              int       `json:"id"`
	IdDentist       int       `json:"id_dentist"  binding:"required"`
	IdPatient       int       `json:"id_patient"  binding:"required"`
	AppointmentDate time.Time `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	Description     string    `json:"description" binding:"required"`
}
