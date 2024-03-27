package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type jsonStore struct {
	pathToFile string
}


func (s *jsonStore) loadPatients() ([]domain.Patient, error) {
	var patients []domain.Patient
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &patients)
	if err != nil {
		return nil, err
	}
	return patients, nil
}

// savePatients guarda los pacientes en un archivo json
func (s *jsonStore) savePatients(patients []domain.Patient) error {
	bytes, err := json.Marshal(patients)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

// NewJsonStore crea un nuevo store de paciente
func NewJsonStore(path string) StoreInterfacePatient {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}

func (s *jsonStore)GetAll() ([]domain.Patient, error) {
	return nil, nil
}


func (s *jsonStore) Read(id int) (domain.Patient, error) {
	patients, err := s.loadPatients()
	if err != nil {
		return domain.Patient{}, err
	}
	for _, patient := range patients {
		if patient.Id == id {
			return patient, nil
		}
	}
	return domain.Patient{}, errors.New("patient not found")
}

func (s *jsonStore) Create(patient domain.Patient) error {
	patients, err := s.loadPatients()
	if err != nil {
		return err
	}
	patient.Id = len(patients) + 1
	patients = append(patients, patient)
	return s.savePatients(patients)
}

func (s *jsonStore) Update(patient domain.Patient) error {
	patients, err := s.loadPatients()
	if err != nil {
		return err
	}
	for i, p := range patients {
		if p.Id == patient.Id {
			patients[i] = patient
			return s.savePatients(patients)
		}
	}
	return errors.New("patient not found")
}

func (s *jsonStore) Delete(id int) error {
	patients, err := s.loadPatients()
	if err != nil {
		return err
	}
	for i, p := range patients {
		if p.Id == id {
			patients = append(patients[:i], patients[i+1:]...)
			return s.savePatients(patients)
		}
	}
	return errors.New("product not found")}

func (s *jsonStore) Exists(dni string) bool {
	patients, err := s.loadPatients()
	if err != nil {
		return false
	}
	for _, p := range patients {
		if p.DNI == dni {
			return true
		}
	}
	return false
}
