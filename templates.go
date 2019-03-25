package main

import (
	"net/http"
	"html/template"
)

type TodoPageData struct {
	PageTitle string
	Todos     []   Todo
}

type Todo struct {
	Title string
	Done  bool
}



func main() {
	tmplate := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		data := TodoPageData{
			PageTitle: "My TODO List",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: true},
			},
		}
		tmplate.Execute(writer, data)
	})

	http.ListenAndServe(":8080", nil)

}



