package todos

import (
	"encoding/json"
	"html/template"
	"net/http"
)

var todos = []Todo{
	{
		Id:    1,
		Title: "Item 1",
		State: INPROGRESS,
	},
	{
		Id:    2,
		Title: "Item 2",
		State: DONE,
	},
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	tmplFile := "todo.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func All(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
