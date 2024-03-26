package appointment

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Repository interface {
	Create(p domain.Appointment) (domain.Appointment, error)
	GetAll() []domain.Appointment
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
