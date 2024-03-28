package dentist

import (
	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Service interface {
	Create(p domain.Dentist) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	GetByID(id int) (domain.Dentist, error)
	Update(id int, d domain.Dentist)(domain.Dentist, error)
	Patch(id int, d domain.Dentist) (domain.Dentist, error)
	Delete(id int) error

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
	l, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
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

//Update dentist
func (s *service) Update(id int, u domain.Dentist) (domain.Dentist, error){
	
	p, err := s.r.GetByID(id)
	if err != nil{
		return domain.Dentist{}, err
	}
	if u.Lastname != ""{
		p.Lastname = u.Lastname
	}
	if u.Name != ""{
		p.Name = u.Name
	}
		if u.License != "" {
		p.License = u.License
	}
		
	p, err = s.r.Update(id, p)

	if err != nil {
		return domain.Dentist{}, err
	}

	return p, nil
}


//Update columns
func (s *service) Patch(id int, u domain.Dentist) (domain.Dentist, error){
	
	p, err := s.r.GetByID(id)
	if err != nil{
		return domain.Dentist{}, err
	}
	if u.Lastname != ""{
		p.Lastname = u.Lastname
	}
	if u.Name != ""{
		p.Name = u.Name
	}
	if u.License != "" {
		p.License = u.License
	}
		
	p, err = s.r.Patch(id, p)

	if err != nil {
		return domain.Dentist{}, err
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