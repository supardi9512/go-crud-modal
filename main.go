package main

import (
	"go-crud-modal/controllers/patientcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", patientcontroller.Index)
	http.HandleFunc("/patient/get_form", patientcontroller.GetForm)
	http.HandleFunc("/patient/store", patientcontroller.Store)

	http.ListenAndServe(":3000", nil)
}
