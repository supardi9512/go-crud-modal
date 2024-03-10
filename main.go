package main

import (
	"go-crud-modal/controllers/studentcontroller"
	"net/http"
)

func main() {
	http.HandleFunc("/", studentcontroller.Index)

	http.ListenAndServe(":8000", nil)
}
