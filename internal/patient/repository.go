package patient

import (
	"errors"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Repository interface {
	Create(p domain.Patient) (domain.Patient, error)
	GetAll() []domain.Patient
	GetByID(id int) (domain.Patient, error)
}

type repository struct {
	list []domain.Patient
}

func NewRepository(list []domain.Patient) Repository {
	return &repository{list}
}

// Create a new patient
func (r *repository) Create(p domain.Patient) (domain.Patient, error) {
	if !r.validateCodeValue(p.DNI) {
		return domain.Patient{}, errors.New("patient already exists")
	}
	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// Validate patient does not exists
func (r *repository) validateCodeValue(dni string) bool {
	for _, patient := range r.list {
		if patient.DNI == dni {
			return false
		}
	}
	return true
}

// GetAll devuelve todos los pacientes
func (r *repository) GetAll() []domain.Patient {
	return r.list
}

// Get patient by ID
func (r *repository) GetByID(id int) (domain.Patient, error) {
	for _, patient := range r.list {
		if patient.Id == id {
			return patient, nil
		}
	}
	return domain.Patient{}, errors.New("patient not found")

}
