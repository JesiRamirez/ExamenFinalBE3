package domain

import (
	 "/appointment"
	)

type Dentist struct {
	Id				int						`json:"id"`
	Lastname		string					`json:"last_name" binding:"required"`
	Name			string					`json:"name" binding:"required"`
	License			string					`json:"license" binding:"required"`
	Appointments 	[]appointment.Appointments		`json:"appointment"`
}