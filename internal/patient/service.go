package patient

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Service interface {
	Create(p domain.Patient) (domain.Patient, error)
}

type service struct {
	r Repository
}

// Create a new service
func NewService(r Repository) Service {
	return &service{r}
}

// Create a new patient
func (s *service) Create(p domain.Patient) (domain.Patient, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}
