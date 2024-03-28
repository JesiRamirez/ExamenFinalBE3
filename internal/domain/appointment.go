package domain

import "time"

type Appointment struct {
	Id          int       `json:"id" binding:"required"`
	PatientId   string       `json:"patient_id" binding:"required"`
	DentistId   string       `json:"dentist_id" binding:"required"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}
