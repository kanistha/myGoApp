package main

import (
	"net/http"
	"html/template"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {

	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			tmpl.Execute(writer, nil)
		}

		contactDetails := ContactDetails{
			Email:   request.FormValue("email"),
			Subject: request.FormValue("subject"),
			Message: request.FormValue("message"),
		}

		_ = contactDetails

		tmpl.Execute(writer, struct {Success bool}{true})
	})
	http.ListenAndServe(":8080", nil)

}