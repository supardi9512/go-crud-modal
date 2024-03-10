package studentcontroller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/student/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
