package appointment

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Service interface {
	Create(p domain.Appointment) (domain.Appointment, error)
	GetAll() ([]domain.Appointment, error)
	GetByID(id int) (domain.Appointment, error)
	Update(id int, p domain.Appointment) (domain.Appointment, error)
	Patch(id int, p domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
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

// Get appointment by ID
func (s *service) GetByID(id int) (domain.Appointment, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return p, nil
}


// Update appointment
func (s *service) Update(id int, u domain.Appointment) (domain.Appointment, error) {

	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	if u.PatientId != ""{
		p.PatientId = u.PatientId
	}
	if u.DentistId != ""{
		p.DentistId = u.DentistId
	}
	if u.Description != "" {
		p.Description = u.Description
	}
	
	p, err = s.r.Update(id, p)

	if err != nil {
		return domain.Appointment{}, err
	}

	return p, nil
}

// Patch appointment
func (s *service) Patch(id int, u domain.Appointment) (domain.Appointment, error) {

	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	if u.PatientId != "" {
		p.PatientId = u.PatientId
	}
	if u.DentistId != "" {
		p.DentistId = u.DentistId
	}
	if u.Description != "" {
		p.Description = u.Description
	}
	
	p, err = s.r.Patch(id, p)

	if err != nil {
		return domain.Appointment{}, err
	}

	return p, nil
}

// Delete appointment
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}