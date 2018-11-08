//Demo using template to write html output

package main

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)


type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}


func main() {

	//load the template from disk
	tmpl := template.Must(template.ParseFiles("template/layout.html"))

	//create the mux router
	r := mux.NewRouter()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My Todo list",
			Todos: []Todo{
				{Title:"task 1", Done:false},
				{Title:"task 2", Done:false},
				{Title:"task 3", Done:true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", r)
}
