package appointment

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Service interface {
	Create(p domain.Appointment) (domain.Appointment, error)
	GetAll() ([]domain.Appointment, error)
}

type service struct {
	r Repository
}

// Create a new service
func NewService(r Repository) Service {
	return &service{r}
}

// Create a new appointment
func (s *service) Create(p domain.Appointment) (domain.Appointment, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Appointment{}, err
	}
	return p, nil
}

// Get all Appointment
func (s *service) GetAll() ([]domain.Appointment, error) {
	l := s.r.GetAll()
	return l, nil
}
