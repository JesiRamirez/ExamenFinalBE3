package patient

import (
	"github.com/JesiRamirez/ExamenFinalBE3.git/internal/domain"
	"net/url"
	
)

type Service interface {
	GetByID(id int) (domain.Patient, error)
	Create(p domain.Patient) (domain.Patient, error)
	Delete(id int) (domain.Patient, error)
	Update(id int, p domain.Patient) (domain.Patient, error)
	GetByParams(searchParams url.Values) ([]domain.Patient, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service(r)
}