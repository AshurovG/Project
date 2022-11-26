package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Login         string
	Password      string
	Success       bool
	StorageAccess string
}

var Users = map[string]string{
	"Ashurov":    "123321123321",
	"Kolomichyk": "3444566666",
}

type LP struct {
	Login, Password string
}

type Validate interface {
	validate() bool
}

func (t *LP) validate() bool {
	for key, el := range Users {
		if key == t.Login && el == t.Password {
			return true
		}
	}
	return false
}

var (
	tmpl = template.Must(template.ParseFiles("forms.html"))
)

func handler(w http.ResponseWriter, req *http.Request) {
	data := ContactDetails{
		Login:    req.FormValue("login"),
		Password: req.FormValue("password"),
	}
	var tmp Validate = LP{data.Login, data.Password}
	data.Success = tmp.validate()
}
