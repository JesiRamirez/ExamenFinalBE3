package appointmentDNILicense

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Service interface {
	Create(p domain.AppointmentDNILicense) (domain.AppointmentDNILicense, error)
	GetByPatientDNI(dni string) (domain.AppointmentDNILicense, error)
}

type service struct {
	r Repository
}

// Create a new service
func NewService(r Repository) Service {
	return &service{r}
}

// Create a new AppointmentDNILicense
func (s *service) Create(p domain.AppointmentDNILicense) (domain.AppointmentDNILicense, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.AppointmentDNILicense{}, err
	}
	return p, nil
}

// Get AppointmentDNILicense by ID
func (s *service) GetByPatientDNI(dni string) (domain.AppointmentDNILicense, error) {
	p, err := s.r.GetByPatientDNI(dni)
	if err != nil {
		return domain.AppointmentDNILicense{}, err
	}
	return p, nil
}
