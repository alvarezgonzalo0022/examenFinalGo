package domain

import "time"

type Paciente struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"address"`
	DNI         string    `json:"dni"`
	IngressDate time.Time  `json:"ingress_date"`
}