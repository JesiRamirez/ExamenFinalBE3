package dentistStore

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type StoreInterfaceDentist interface {
	//Read devuelve un dentista por su id
	Read(id int) (domain.Dentist, error)
	//Create agrega un dentista nuevo
	Create(dentist domain.Dentist) error
	//Update actualiza un dentista
	Update(dentist domain.Dentist) error
	//Delete elimina un dentista
	Delete(id int) error
	//Exists verifica si un dentista
	Exists(license string) bool
	//GetAll obtiene todos los dentistas
	GetAll() ([]domain.Dentist, error)

}