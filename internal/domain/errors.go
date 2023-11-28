package domain

import "errors"

var (
	ErrPrepareStatement    = errors.New("error prepare statement")
	ErrExecStatement       = errors.New("error exec statement")
	ErrLastInsertedId      = errors.New("error last inserted id")
	ErrEmpty               = errors.New("empty list")
	ErrPatientNotFound     = errors.New("patient not found")
	ErrDentistNotFound     = errors.New("dentist not found")
	ErrAppointmentNotFound = errors.New("appointment not found")
	ErrInvalidID           = errors.New("error invalid id")
	ErrScanRow             = errors.New("error scanning row")
	ErrQuery			   = errors.New("error query")
)
