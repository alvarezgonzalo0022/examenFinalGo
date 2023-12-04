package domain

import "time"

type Patient struct {
	Id            int       `json:"id"`
	FirstName     string    `json:"first_name" binding:"required"`
	LastName      string    `json:"last_name" binding:"required"`
	Address       string    `json:"address"`
	Document      string    `json:"document" binding:"required"`
	AdmissionDate time.Time `json:"admission_date" binding:"required"`
}
