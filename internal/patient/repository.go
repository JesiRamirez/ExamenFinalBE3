package patient

import (
	"errors"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/store"
)

type Repository interface {
	Create(p domain.Patient) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	GetByID(id int) (domain.Patient, error)
	Update(id int, p domain.Patient) (domain.Patient, error)
	Patch(id int, p domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfacePatient
}

func NewRepository(storage store.StoreInterfacePatient) Repository {
	return &repository{storage}
}

// Create a new patient
func (r *repository) Create(p domain.Patient) (domain.Patient, error) {
	if !r.storage.Exists(p.DNI) {
		return domain.Patient{}, errors.New("patient already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return	domain.Patient{}, errors.New("error creating patitent")
	}
	return p, nil
}


// GetAll devuelve todos los pacientes
func (r *repository) GetAll() ([]domain.Patient, error) {
	patients, err := r.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return patients, nil
}

// Get patient by ID
func (r *repository) GetByID(id int) (domain.Patient, error) {
	patient, err := r.storage.Read(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil

}

// Update patient
func (r *repository) Update(id int, p domain.Patient) (domain.Patient, error) {

	if !r.storage.Exists(p.DNI) {
		return domain.Patient{}, errors.New("dni already exists")
	}
	err := r.storage.Update(p)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return p, nil
}

// Patch patient
func (r *repository) Patch(id int, p domain.Patient) (domain.Patient, error) {

	err := r.storage.Update(p)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return p, nil
}

// Delete patient
func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
