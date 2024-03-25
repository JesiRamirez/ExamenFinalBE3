package dentist

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Service interface {
	Create(p domain.Dentist) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	GetByID(id int) (domain.Dentist, error)
}

type service struct {
	r Repository
}

// Create a new service
func NewService(r Repository) Service {
	return &service{r}
}

// Create a new dentist
func (s *service) Create(p domain.Dentist) (domain.Dentist, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

// Get all dentists
func (s *service) GetAll() ([]domain.Dentist, error) {
	l := s.r.GetAll()
	return l, nil
}

// Get dentists by ID
func (s *service) GetByID(id int) (domain.Dentist, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}
