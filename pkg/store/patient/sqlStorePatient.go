package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
)

type sqlStorePatient struct {
	db *sql.DB
}

func NewSqlStorePatient(db *sql.DB) StoreInterfacePatient {
	return &sqlStorePatient{
		db: db,
	}
}

func (s *sqlStorePatient) Read(id int) (domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patients WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&patient.Id, &patient.Lastname, &patient.Name, &patient.Address, &patient.DNI, &patient.AdmissionDate)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *sqlStorePatient) GetAll() ([]domain.Patient, error) {
	listReturn := []domain.Patient{}
	query := "SELECT * FROM patients;"

	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var patient domain.Patient
		err := rows.Scan(&patient.Id, &patient.Lastname, &patient.Name, &patient.Address, &patient.DNI, &patient.AdmissionDate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(patient)
		listReturn = append(listReturn, patient)
	}

	return listReturn, nil
}

// Create agrega un nuevo paciente
func (s *sqlStorePatient) Create(patient domain.Patient) error {
	query := "INSERT INTO patients (last_name, name, adress, dni, admission_date) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(patient.Lastname, patient.Name, patient.Address, patient.DNI, patient.AdmissionDate)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Update actualiza un paciente
func (s *sqlStorePatient) Update(patient domain.Patient) error {
	query := "UPDATE patients SET last_name = ?, name = ?, adress = ?, dni = ?, admission_date = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(patient.Lastname, patient.Name, patient.Address, patient.DNI, patient.AdmissionDate, patient.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// Delete elimina un paciente
func (s *sqlStorePatient) Delete(id int) error {
	query := "DELETE FROM patients WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// Exists verifica si un paciente existe
func (s *sqlStorePatient) Exists(dni string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM patients WHERE dni = ?;"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
