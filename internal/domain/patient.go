package domain

import "time"

type Patient struct {
	Id            int       `json:"id"`
	Lastname      string    `json:"last_name" binding:"required"`
	Name          string    `json:"name" binding:"required"`
	Address       string    `json:"adress"`
	DNI           string    `json:"dni"`
	AdmissionDate time.Time `json:"admission_date"`
	//Appointments	[]appointment.Appointment		`json:"appointment" binding:"required"`
}
