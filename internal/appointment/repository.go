package appointment

import (
	"errors"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/store/AppointmentStore"
)

type Repository interface {
	Create(p domain.Appointment) (domain.Appointment, error)
	GetAll() ([]domain.Appointment, error)
	GetByID(id int) (domain.Appointment, error)
	Update(id int, p domain.Appointment) (domain.Appointment, error)
	Patch(id int, p domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type repository struct {
	storageAppointment appointmentStore.StoreInterfaceAppointment
}

func NewRepository(storageAppointment appointmentStore.StoreInterfaceAppointment) Repository {
	return &repository{storageAppointment}
}

// Create a new appointment
func (r *repository) Create(p domain.Appointment) (domain.Appointment, error) {

	if !r.storageAppointment.Exists(p.Id) {
		return domain.Appointment{}, errors.New("id already exists")
	}
	err := r.storageAppointment.Create(p)
	if err != nil {
		return	domain.Appointment{}, errors.New("error creating appointment")
	}
	return p, nil
}

// GetAll devuelve todos los appointment
func (r *repository) GetAll() ([]domain.Appointment, error) {
	appointments, err := r.storageAppointment.GetAll()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

// Get appointment by ID
func (r *repository) GetByID(id int) (domain.Appointment, error) {
	appointment, err := r.storageAppointment.Read(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}
	return appointment, nil

}

// Update Appointment
func (r *repository) Update(id int, p domain.Appointment) (domain.Appointment, error) {

	if !r.storageAppointment.Exists(p.Id) {
		return domain.Appointment{}, errors.New("id already exists")
	}
	err := r.storageAppointment.Update(p)
	if err != nil {
		return domain.Appointment{}, errors.New("error updating appointment")
	}
	return p, nil
}

// Patch patient
func (r *repository) Patch(id int, p domain.Appointment) (domain.Appointment, error) {

	err := r.storageAppointment.Update(p)
	if err != nil {
		return domain.Appointment{}, errors.New("error updating appointment")
	}
	return p, nil
}

// Delete appointment
func (r *repository) Delete(id int) error {
	err := r.storageAppointment.Delete(id)
	if err != nil {
		return err
	}
	return nil
}