package patient

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Service interface {
	Create(p domain.Patient) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	GetByID(id int) (domain.Patient, error)
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

// Get all patients
func (s *service) GetAll() ([]domain.Patient, error) {
	l := s.r.GetAll()
	return l, nil
}

// Get patients by ID
func (s *service) GetByID(id int) (domain.Patient, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}
