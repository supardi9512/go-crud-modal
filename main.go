package main

import (
	"go-crud-modal/controllers/patientcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", patientcontroller.Index)

	http.ListenAndServe(":3000", nil)
}
