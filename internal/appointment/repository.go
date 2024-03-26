package appointment

import (
	"errors"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Repository interface {
	Create(p domain.Appointment) (domain.Appointment, error)
	GetAll() []domain.Appointment
	GetByID(id int) (domain.Appointment, error)
	Update(id int, p domain.Appointment) (domain.Appointment, error)
	Patch(id int, p domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type repository struct {
	list []domain.Appointment
}

func NewRepository(list []domain.Appointment) Repository {
	return &repository{list}
}

// Create a new appointment
func (r *repository) Create(p domain.Appointment) (domain.Appointment, error) {

	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// GetAll devuelve todos los appointment
func (r *repository) GetAll() []domain.Appointment {
	return r.list
}

// Get appointment by ID
func (r *repository) GetByID(id int) (domain.Appointment, error) {
	for _, appointment := range r.list {
		if appointment.Id == id {
			return appointment, nil
		}
	}
	return domain.Appointment{}, errors.New("appointment not found")

}

// Update Appointment
func (r *repository) Update(id int, p domain.Appointment) (domain.Appointment, error) {

	update := false
	for i, v := range r.list {
		if v.Id == id {
			p.Id = id
			r.list[i] = p
			update = true
		}
	}
	if !update {
		return domain.Appointment{}, errors.New("patient not found")
	}
	return p, nil
}

// Patch patient
func (r *repository) Patch(id int, p domain.Appointment) (domain.Appointment, error) {

	update := false
	for i, v := range r.list {
		if v.Id == id {
			p.Id = id
			r.list[i] = p
			update = true
		}
	}
	if !update {
		return domain.Appointment{}, errors.New("appointment not found")
	}
	return p, nil
}

// Delete appointment
func (r *repository) Delete(id int) error {
	deleted := false
	for i, v := range r.list {
		if v.Id == id {
			r.list = append(r.list[:i], r.list[i+1:]...)
			deleted = true
		}
	}
	if !deleted {
		return errors.New("appointment not found")
	}
	return nil
}