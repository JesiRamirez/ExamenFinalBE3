package appointmentDNILicense

import (
	"errors"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type Repository interface {
	Create(p domain.AppointmentDNILicense) (domain.AppointmentDNILicense, error)
	GetByPatientDNI(dni string) (domain.AppointmentDNILicense, error)
}

type repository struct {
	list []domain.AppointmentDNILicense
}

func NewRepository(list []domain.AppointmentDNILicense) Repository {
	return &repository{list}
}

// Create a new AppointmentDNILicense
func (r *repository) Create(p domain.AppointmentDNILicense) (domain.AppointmentDNILicense, error) {

	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// GetAll devuelve todos los AppointmentDNILicense
func (r *repository) GetAll() []domain.AppointmentDNILicense {
	return r.list
}

// Get AppointmentDNILicense by ID
func (r *repository) GetByPatientDNI(dni string) (domain.AppointmentDNILicense, error) {
	for _, appointmentDNILicense := range r.list {
		if appointmentDNILicense.PatientDNI == dni {
			return appointmentDNILicense, nil
		}
	}
	return domain.AppointmentDNILicense{}, errors.New("AppointmentDNILicense not found")

}
