package patient

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Service interface {
	Create(p domain.Patient) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	GetByID(id int) (domain.Patient, error)
	Update(id int, p domain.Patient)(domain.Patient, error)
	Patch(id int, p domain.Patient) (domain.Patient, error)
	Delete(id int) error
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

//Update patient
func (s *service) Update(id int, u domain.Patient) (domain.Patient, error){
	
	p, err := s.r.GetByID(id)
	if err != nil{
		return domain.Patient{}, err
	}
	if u.Lastname != ""{
		p.Lastname = u.Lastname
	}
	if u.Name != ""{
		p.Name = u.Name
	}
	if u.Address != ""{
		p.Address = u.Address
	}
	if u.DNI != "" {
		p.DNI = u.DNI
	}
		
	p, err = s.r.Update(id, p)

	if err != nil {
		return domain.Patient{}, err
	}

	return p, nil
}


//Update columns
func (s *service) Patch(id int, u domain.Patient) (domain.Patient, error){
	
	p, err := s.r.GetByID(id)
	if err != nil{
		return domain.Patient{}, err
	}
	if u.Lastname != ""{
		p.Lastname = u.Lastname
	}
	if u.Name != ""{
		p.Name = u.Name
	}
	if u.Address != ""{
		p.Address = u.Address
	}
	if u.DNI != "" {
		p.DNI = u.DNI
	}
		
	p, err = s.r.Patch(id, p)

	if err != nil {
		return domain.Patient{}, err
	}

	return p, nil
}

//Delete patient
func (s *service) Delete(id int) error{
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}