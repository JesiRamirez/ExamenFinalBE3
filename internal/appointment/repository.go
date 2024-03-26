package appointment

import (
	"errors"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Repository interface {
	Create(p domain.Appointment) (domain.Appointment, error)
	GetAll() []domain.Appointment
	GetByID(id int) (domain.Appointment, error)
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
