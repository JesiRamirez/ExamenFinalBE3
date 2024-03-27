package appointmentStore


import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type jsonStoreAppointment struct {
	pathToFile string
}


func (s *jsonStoreAppointment) loadAppointments() ([]domain.Appointment, error) {
	var appointment []domain.Appointment
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &appointment)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *jsonStoreAppointment) saveAppointments(appointment []domain.Appointment) error {
	bytes, err := json.Marshal(appointment)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

func NewjsonStoreAppointment(path string) StoreInterfaceAppointment {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStoreAppointment{  
		pathToFile: path,
	}
}

func (s *jsonStoreAppointment)GetAll() ([]domain.Appointment, error) {
	return nil, nil
}


func (s *jsonStoreAppointment) Read(id int) (domain.Appointment, error) {
	appointments, err := s.loadAppointments()
	if err != nil {
		return domain.Appointment{}, err
	}
	for _, appointment := range appointments {
		if appointment.Id == id {
			return appointment, nil
		}
	}
	return domain.Appointment{}, errors.New("appointment not found")
}

func (s *jsonStoreAppointment) Create(appointment domain.Appointment) error {
	appointments, err := s.loadAppointments()
	if err != nil {
		return err
	}
	appointment.Id = len(appointments) + 1
	appointments = append(appointments, appointment)
	return s.saveAppointments(appointments)
}

func (s *jsonStoreAppointment) Update(appointment domain.Appointment) error {
	appointments, err := s.loadAppointments()
	if err != nil {
		return err
	}
	for i, p := range appointments {
		if p.Id == appointment.Id {
			appointments[i] = appointment
			return s.saveAppointments(appointments)
		}
	}
	return errors.New("appointment not found")
}

func (s *jsonStoreAppointment) Delete(id int) error {
	appointments, err := s.loadAppointments()
	if err != nil {
		return err
	}
	for i, p := range appointments {
		if p.Id == id {
			appointments = append(appointments[:i], appointments[i+1:]...)
			return s.saveAppointments(appointments)
		}
	}
	return errors.New("appointment not found")}

func (s *jsonStoreAppointment) Exists(id int) bool {
	appointments, err := s.loadAppointments()
	if err != nil {
		return false
	}
	for _, p := range appointments {
		if p.Id == id {
			return true
		}
	}
	return false
}
