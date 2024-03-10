package patientmodel

import (
	"database/sql"
	"go-crud-modal/config"
	"go-crud-modal/entities"
	"time"
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

		if data.Gender == "1" {
			data.Gender = "Male"
		} else if data.Gender == "2" {
			data.Gender = "Female"
		}

		// 2006-01-02 => yyyy-mm-dd
		date_of_birth, _ := time.Parse("2006-01-02", data.DateOfBirth)

		// 02-01-2006 => dd-mm-yyyy
		data.DateOfBirth = date_of_birth.Format("02-01-2006")

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

func (p *PatientModel) Find(id int64, patient *entities.Patient) error {
	return p.db.QueryRow("SELECT * FROM patients WHERE id = ?", id).Scan(&patient.Id, &patient.Name, &patient.Nik, &patient.Gender, &patient.PlaceOfBirth, &patient.DateOfBirth, &patient.Address, &patient.PhoneNumber)
}

func (p *PatientModel) Update(patient entities.Patient) error {
	_, err := p.db.Exec("UPDATE patients SET name = ?, nik = ?, gender = ?, place_of_birth = ?, date_of_birth = ?, address = ?, phone_number = ? WHERE id = ?", patient.Name, patient.Nik, patient.Gender, patient.PlaceOfBirth, patient.DateOfBirth, patient.Address, patient.PhoneNumber, patient.Id)

	if err != nil {
		return err
	}

	return nil
}
