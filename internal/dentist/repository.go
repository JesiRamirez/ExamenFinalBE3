package dentist

import (
	"errors"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/store/dentistStore"
)

type Repository interface {
	Create(p domain.Dentist) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	GetByID(id int) (domain.Dentist, error)
	Update(id int, p domain.Dentist) (domain.Dentist, error)
	Patch(id int, p domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type repository struct {
	storageDentist dentistStore.StoreInterfaceDentist
}

func NewRepository(storageDentist dentistStore.StoreInterfaceDentist) Repository {
	return &repository{storageDentist}
}

// Create a new dentist
func (r *repository) Create(p domain.Dentist) (domain.Dentist, error) {
	if !r.storageDentist.Exists(p.License) {
		return domain.Dentist{}, errors.New("dentist already exists")
	}
	err := r.storageDentist.Create(p)
	if err != nil {
		return	domain.Dentist{}, errors.New("error creating dentist")
	}
	return p, nil
}


// GetAll dentists
func (r *repository) GetAll() ([]domain.Dentist, error) {
	dentists, err := r.storageDentist.GetAll()
	if err != nil {
		return nil, err
	}
	return dentists, nil}

// Get dentist by ID
func (r *repository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.storageDentist.Read(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	return dentist, nil

}

//Update Dentist
func (r *repository) Update(id int, p domain.Dentist) (domain.Dentist, error){
	if !r.storageDentist.Exists(p.License) {
		return domain.Dentist{}, errors.New("license already exists")
	}
	err := r.storageDentist.Update(p)
	if err != nil {
		return domain.Dentist{}, errors.New("error updating dentist")
	}
	return p, nil
}

//Patch dentist
func (r *repository) Patch(id int, p domain.Dentist) (domain.Dentist, error){
	
	err := r.storageDentist.Update(p)
	if err != nil {
		return domain.Dentist{}, errors.New("error updating dentist")
	}
	return p, nil
}

//Delete dentist
func(r *repository) Delete(id int) error {
	err := r.storageDentist.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
