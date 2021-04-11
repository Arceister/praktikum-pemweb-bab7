package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("View/*.html"))
}

func main() {
	http.HandleFunc("/", ShowProker)
	http.HandleFunc("/create", createProkerPage)
	http.HandleFunc("/edit", editDeleteProkerPage)
	http.HandleFunc("/createproker", createProkerQuery)
	http.HandleFunc("/editproker", editProkerQuery)
	http.HandleFunc("/deleteproker", deleteProkerQuery)
	http.HandleFunc("/login", userLogin)
	http.HandleFunc("/loginprocess", loginUser)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":14022", nil)
}
