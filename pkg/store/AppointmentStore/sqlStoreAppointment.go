package appointmentStore

import (
	"database/sql"
	"log"
	"fmt"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	
)

type sqlStoreAppointment struct {
	db *sql.DB
}

func NewSqlStoreAppointment(db *sql.DB) StoreInterfaceAppointment {
	return &sqlStoreAppointment{
		db: db,
	}
}

func (s *sqlStoreAppointment) Read(id int) (domain.Appointment, error) {
	var appointment domain.Appointment
	query := "SELECT * FROM appointments WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&appointment.Id, &appointment.PatientId, &appointment.DentistId, &appointment.Date, &appointment.Description)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *sqlStoreAppointment) GetAll() ([]domain.Appointment, error) {
	listReturn := []domain.Appointment{}
	query := "SELECT * FROM appointments;"

	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(&appointment.Id, &appointment.PatientId, &appointment.DentistId, &appointment.Date, &appointment.Description)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(appointment)
		listReturn = append(listReturn, appointment)
	}

	return listReturn, nil
}

//Create agrega un nuevo turno
func (s *sqlStoreAppointment) Create(appointment domain.Appointment) error {
	query := "INSERT INTO appointments (patient_id, dentist_id, date, description) VALUES (?, ?, ?, ? );"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	
	res, err := stmt.Exec(appointment.PatientId, appointment.DentistId, appointment.Date, appointment.Description)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

//Update actualiza un turno
func (s *sqlStoreAppointment) Update(appointment domain.Appointment) error {
	query := "UPDATE appointments SET patient_id = ?, dentist_id = ?, date = ?, description = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(appointment.PatientId, appointment.DentistId, appointment.Date, appointment.Description, appointment.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

//Delete elimina un turno
func (s *sqlStoreAppointment) Delete(id int) error {
	query := "DELETE FROM appointments WHERE id = ?;"
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

//Exists verifica si un turno existe
func (s *sqlStoreAppointment) Exists(id int) bool {
	var exists bool
	var id1 int
	query := "SELECT id FROM dentists WHERE dni = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id1 > 0 {
		exists = true
	}
	return exists
}
