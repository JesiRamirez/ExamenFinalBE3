package patient

import (
	"/DesafioFinal/internal/domain"
	"/DesafioFinal/pkg/store"

	"errors"
	"net/url"
)

type Repository interface {

	GetByID(id int) (domain.Patient, error)

	Create(p domain.Patient) (domain.Patient, error)

	Update(id int, p domain.Patient) (domain.Patient, error)

	Delete(id int) error
}

type repository struct{
	storage store.StoreInterface 
}

//NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) Repository{
	return &repository(storage)
}

func (r *repository) GetByID(id int) (domain.Patient, error){
	patient, err := r.storage.Read(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil
}