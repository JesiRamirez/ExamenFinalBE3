package domain

import "time"

type appointment struct {
	Id        	int 		`json:"id" binding:"required"`
	PatientId 	int			`json:"pacient_id" binding:"required"`
	DentistId 	int			`json:"dentist_id" binding:"required"`
	Date      	time.Time	`json:"date"`
	Description	string		`json:"description"`
}