package main

import "net/http"

func createProkerPage(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized {
		tpl.ExecuteTemplate(w, "notauthorized.html", nil)
	} else {
		tpl.ExecuteTemplate(w, "create.html", nil)
	}
}

func editDeleteProkerPage(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized {
		tpl.ExecuteTemplate(w, "notauthorized.html", nil)
	} else {
		tpl.ExecuteTemplate(w, "editdelete.html", nil)
	}
}

func ShowProker(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized {
		tpl.ExecuteTemplate(w, "notauthorized.html", nil)
	} else {
		AllProker := returnProker()

		allProker := SemuaProker{
			AllProker,
		}

		tpl.ExecuteTemplate(w, "index.html", allProker)
	}
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	isAuthorized = false
	tpl.ExecuteTemplate(w, "return.html", nil)
}
