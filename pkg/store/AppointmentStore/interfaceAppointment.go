package appointmentStore

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type StoreInterfaceAppointment interface {
	//Read devuelve un turno por su id
	Read(id int) (domain.Appointment, error)
	//Create agrega un turno nuevo
	Create(dentist domain.Appointment) error
	//Update actualiza un turno
	Update(dentist domain.Appointment) error
	//Delete elimina un turno
	Delete(id int) error
	//Exists verifica si un turno
	Exists(id int) bool
	//GetAll obtiene todos los turno
	GetAll() ([]domain.Appointment, error)

}