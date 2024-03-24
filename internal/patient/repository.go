package repository

import (
	"\DesafioFinal\internal\domain"
	
)

type Repository interface {

	GetByID(id int) (domain.Patient, error)

	Create(p domain.Patient) (domain.Patient, error)

	Update(id int, p domain.Patient) (domain.Patient, error)

	Delete(id int) error
}

type repository struct{
	storage StoreInterface 
}