package domain

import "time"

type AppointmentDNILicense struct {
	Id             int       `json:"id" binding:"required"`
	PatientDNI     string    `json:"pacient_dni" binding:"required"`
	DentistLicense string    `json:"dentist_license" binding:"required"`
	Date           time.Time `json:"date"`
	Description    string    `json:"description"`
}
