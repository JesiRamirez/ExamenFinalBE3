package dentistStore

import (
	"database/sql"
	"log"
	"fmt"

	"github.com/bootcamp-go/ExamenFinalBE3.git/internal/domain"
	
)

type sqlStoreDentist struct {
	db *sql.DB
}

func NewSqlStoreDentist(db *sql.DB) StoreInterfaceDentist {
	return &sqlStoreDentist{
		db: db,
	}
}

func (s *sqlStoreDentist) Read(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	query := "SELECT * FROM dentists WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&dentist.Id, &dentist.Lastname, &dentist.Name, &dentist.License)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *sqlStoreDentist) GetAll() ([]domain.Dentist, error) {
	listReturn := []domain.Dentist{}
	query := "SELECT * FROM dentists;"

	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var dentist domain.Dentist
		err := rows.Scan(&dentist.Id, &dentist.Lastname, &dentist.Name, &dentist.License)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dentist)
		listReturn = append(listReturn, dentist)
	}

	return listReturn, nil
}

//Create agrega un nuevo dentista
func (s *sqlStoreDentist) Create(dentist domain.Dentist) error {
	query := "INSERT INTO dentists (last_name, name, license) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	
	res, err := stmt.Exec(dentist.Lastname, dentist.Name, dentist.License)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

//Update actualiza un dentista
func (s *sqlStoreDentist) Update(dentist domain.Dentist) error {
	query := "UPDATE dentist SET last_name = ?, name = ?, license = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentist.Lastname, dentist.Name, dentist.License, dentist.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

//Delete elimina un dentist
func (s *sqlStoreDentist) Delete(id int) error {
	query := "DELETE FROM dentists WHERE id = ?;"
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

//Exists verifica si un dentist existe
func (s *sqlStoreDentist) Exists(license string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM dentists WHERE dni = ?;"
	row := s.db.QueryRow(query, license)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
