package dentist

import (
	"errors"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Repository interface {
	Create(p domain.Dentist) (domain.Dentist, error)
	GetAll() []domain.Dentist
	GetByID(id int) (domain.Dentist, error)
	Update(id int, p domain.Dentist) (domain.Dentist, error)
	Patch(id int, p domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type repository struct {
	list []domain.Dentist
}

func NewRepository(list []domain.Dentist) Repository {
	return &repository{list}
}

// Create a new dentist
func (r *repository) Create(p domain.Dentist) (domain.Dentist, error) {
	if !r.validateLicense(p.License) {
		return domain.Dentist{}, errors.New("dentist already exists")
	}
	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// Validate dentist does not exists
func (r *repository) validateLicense(license string) bool {
	for _, patient := range r.list {
		if patient.License == license {
			return false
		}
	}
	return true
}

// GetAll dentists
func (r *repository) GetAll() []domain.Dentist {
	return r.list
}

// Get dentist by ID
func (r *repository) GetByID(id int) (domain.Dentist, error) {
	for _, dentist := range r.list {
		if dentist.Id == id {
			return dentist, nil
		}
	}
	return domain.Dentist{}, errors.New("dentist not found")

}

//Update Dentist
func (r *repository) Update(id int, p domain.Dentist) (domain.Dentist, error){
	if !r.validateLicense(p.License){
		return domain.Dentist{}, errors.New("code value already exists")
	}
	update := false
	for i, v := range r.list {
		if v.Id == id {
			p.Id = id
			r.list[i] = p
			update = true
		}
	}
	if !update {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	return p, nil
}

//Patch dentist
func (r *repository) Patch(id int, p domain.Dentist) (domain.Dentist, error){
	
	update := false
	for i, v := range r.list {
		if v.Id == id {
			p.Id = id
			r.list[i] = p
			update = true
		}
	}
	if !update {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	return p, nil
}

//Delete dentist
func(r *repository) Delete(id int) error {
	deleted := false
	for i, v := range r.list {
		if v.Id == id {
			r.list = append(r.list[:i], r.list[i+1:]...)
			deleted = true
		}
	}
	if !deleted {
		return errors.New("dentist not found")
	}

	return nil
}
