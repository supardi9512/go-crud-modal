package patientcontroller

import (
	"bytes"
	"encoding/json"
	"go-crud-modal/entities"
	"go-crud-modal/models/patientmodel"
	"html/template"
	"net/http"
	"strconv"
)

var patientModel = patientmodel.New()

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/patient/index.html")
	temp.Execute(w, data)
}

func GetData() string {
	buffer := &bytes.Buffer{}

	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/patient/data.html")

	var patient []entities.Patient
	err := patientModel.FindAll(&patient)

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"patient": patient,
	}

	temp.ExecuteTemplate(buffer, "data.html", data)

	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}

	if err != nil {
		data = map[string]interface{}{
			"title": "Add Data",
		}
	} else {
		var patient entities.Patient
		err := patientModel.Find(id, &patient)

		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"title":   "Edit Data",
			"patient": patient,
		}
	}

	temp, _ := template.ParseFiles("views/patient/form.html")
	temp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		var patient entities.Patient

		patient.Name = r.Form.Get("name")
		patient.Nik = r.Form.Get("nik")
		patient.Gender = r.Form.Get("gender")
		patient.PlaceOfBirth = r.Form.Get("place_of_birth")
		patient.DateOfBirth = r.Form.Get("date_of_birth")
		patient.Address = r.Form.Get("address")
		patient.PhoneNumber = r.Form.Get("phone_number")

		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		var data map[string]interface{}

		if err != nil {
			// insert data

			err := patientModel.Create(&patient)

			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}

			data = map[string]interface{}{
				"message": "Patient data has been added successfully",
				"data":    template.HTML(GetData()),
			}
		} else {
			// update data

			patient.Id = id

			err := patientModel.Update(patient)

			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}

			data = map[string]interface{}{
				"message": "Patient data has been updated successfully",
				"data":    template.HTML(GetData()),
			}
		}

		ResponseJson(w, http.StatusOK, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

	if err != nil {
		panic(err)
	}

	err = patientModel.Delete(id)

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"message": "Patient data has been deleted successfully",
		"data":    template.HTML(GetData()),
	}

	ResponseJson(w, http.StatusOK, data)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{
		"error": message,
	})
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
