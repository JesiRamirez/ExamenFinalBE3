package dentistStore

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type jsonStoreDentist struct {
	pathToFile string
}

func (s *jsonStoreDentist) loadDentists() ([]domain.Dentist, error) {
	var dentists []domain.Dentist
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &dentists)
	if err != nil {
		return nil, err
	}
	return dentists, nil
}

func (s *jsonStoreDentist) saveDentists(dentists []domain.Dentist) error {
	bytes, err := json.Marshal(dentists)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

func NewJsonStoreDentist(path string) StoreInterfaceDentist {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStoreDentist{
		pathToFile: path,
	}
}

func (s *jsonStoreDentist) GetAll() ([]domain.Dentist, error) {
	return nil, nil
}

func (s *jsonStoreDentist) Read(id int) (domain.Dentist, error) {
	dentists, err := s.loadDentists()
	if err != nil {
		return domain.Dentist{}, err
	}
	for _, dentist := range dentists {
		if dentist.Id == id {
			return dentist, nil
		}
	}
	return domain.Dentist{}, errors.New("dentist not found")
}

func (s *jsonStoreDentist) Create(dentist domain.Dentist) error {
	dentists, err := s.loadDentists()
	if err != nil {
		return err
	}
	dentist.Id = len(dentists) + 1
	dentists = append(dentists, dentist)
	return s.saveDentists(dentists)
}

func (s *jsonStoreDentist) Update(dentist domain.Dentist) error {
	dentists, err := s.loadDentists()
	if err != nil {
		return err
	}
	for i, p := range dentists {
		if p.Id == dentist.Id {
			dentists[i] = dentist
			return s.saveDentists(dentists)
		}
	}
	return errors.New("dentist not found")
}

func (s *jsonStoreDentist) Delete(id int) error {
	dentists, err := s.loadDentists()
	if err != nil {
		return err
	}
	for i, p := range dentists {
		if p.Id == id {
			dentists = append(dentists[:i], dentists[i+1:]...)
			return s.saveDentists(dentists)
		}
	}
	return errors.New("dentist not found")
}

func (s *jsonStoreDentist) Exists(license string) bool {
	dentists, err := s.loadDentists()
	if err != nil {
		return false
	}
	for _, p := range dentists {
		if p.License == license {
			return true
		}
	}
	return false
}
