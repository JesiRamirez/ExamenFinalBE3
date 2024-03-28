package store

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type StoreInterfacePatient interface {
	//Read devuelve un paciente por su id
	Read(id int) (domain.Patient, error)
	//Create agrega un paciente nuevo
	Create(patient domain.Patient) error
	//Update actualiza un paciente
	Update(patient domain.Patient) error
	//Delete elimina un paciente
	Delete(id int) error
	//Exists verifica si un paciente
	Exists(dni string) bool
	//GetAll obtiene todos los pacientes
	GetAll() ([]domain.Patient, error)

}