package domain

type Dentist struct {
	Id             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	RegistrationId string `json:"registration_id"`
}
