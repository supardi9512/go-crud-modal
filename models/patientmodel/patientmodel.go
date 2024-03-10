package patientmodel

import (
	"database/sql"
	"go-crud-modal/config"
	"go-crud-modal/entities"
)

type PatientModel struct {
	db *sql.DB
}

func New() *PatientModel {
	db, err := config.DBConnection()

	if err != nil {
		panic(err)
	}

	return &PatientModel{
		db: db,
	}
}

func (p *PatientModel) FindAll(patient *[]entities.Patient) error {
	rows, err := p.db.Query("SELECT * FROM patients")

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Patient
		rows.Scan(&data.Id, &data.Name, &data.Nik, &data.Gender, &data.PlaceOfBirth, &data.DateOfBirth, &data.Address, &data.PhoneNumber)

		*patient = append(*patient, data)
	}

	return nil
}

func (p *PatientModel) Create(patient *entities.Patient) error {
	result, err := p.db.Exec("INSERT INTO patients (name, nik, gender, place_of_birth, date_of_birth, address, phone_number) VALUES(?,?,?,?,?,?,?)", patient.Name, patient.Nik, patient.Gender, patient.PlaceOfBirth, patient.DateOfBirth, patient.Address, patient.PhoneNumber)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	patient.Id = lastInsertId

	return nil
}
