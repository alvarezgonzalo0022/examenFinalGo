package domain

import (
	"time"
)

type Appointment struct {
	Id              int 		`json:"id"`
	Patient     	Patient		`json:"patient"  binding:"required"`
	Dentist      	Dentist 	`json:"dentist"  binding:"required"`
	AppointmentDate time.Time 	`json:"appointment_date"`
	AppointmentTime time.Time 	`json:"appointment_time"`
	Description     string		`json:"description" binding:"required"`
}

type AppointmentRequest struct {
	Id              int       	`json:"id"`
	PatientId       int       	`json:"patient_id"  binding:"required"`
	DentistId       int       	`json:"dentist_id"  binding:"required"`
	AppointmentDate string		`json:"appointment_date"`
	AppointmentTime string		`json:"appointment_time"`
	Description     string		`json:"description" binding:"required"`
}

type AppointmentPatchRequest struct {
	Id              int       	`json:"id"`
	PatientId       int       	`json:"patient_id"`
	DentistId       int       	`json:"dentist_id"`
	AppointmentDate string		`json:"appointment_date"`
	AppointmentTime string		`json:"appointment_time"`
	Description     string		`json:"description"`
}

type AppointmentResponse struct {
	Id              int			`json:"id"`
	PatientId       int			`json:"id_patient"  binding:"required"`
	DentistId       int			`json:"id_dentist"  binding:"required"`
	AppointmentDate string		`json:"appointment_date"`
	AppointmentTime string		`json:"appointment_time"`
	Description     string		`json:"description" binding:"required"`
	PatientLastName string		`json:"patient_lastname"`
	DentistLastName string		`json:"dentist_lastname"`
}