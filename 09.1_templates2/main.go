package main

import (
	"net/http"
	"text/template"
)

// Create global variable thats of type pointer to template
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// init

// assign tpl to parseglob showing the path

// create index func and eecute , 3 variables

// handlefunc for index

// start server
