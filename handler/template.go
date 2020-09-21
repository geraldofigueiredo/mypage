package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

type Template struct {
	Filename string
	Templ    *template.Template
}

func (t *Template) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	t.Templ, err = template.ParseFiles(filepath.Join("template", t.Filename))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	err = t.Templ.Execute(w, r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
}
