package patientcontroller

import (
	"bytes"
	"go-crud-modal/entities"
	"go-crud-modal/models/patientmodel"
	"html/template"
	"net/http"
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
	temp, _ := template.ParseFiles("views/patient/form.html")
	temp.Execute(w, nil)
}
