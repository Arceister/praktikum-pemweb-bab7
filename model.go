package main

import (
	"fmt"
	"log"
	"net/http"
)

var isAuthorized bool = false

type Proker struct {
	NomorProgram    int
	NamaProgram     string
	SuratKeterangan string
}

type SemuaProker struct {
	Prokers []Proker
}

func returnProker() []Proker {
	var proker Proker
	var arrProker []Proker

	db := connect()
	rows, err := db.Query("SELECT * FROM proker")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&proker.NomorProgram,
			&proker.NamaProgram,
			&proker.SuratKeterangan,
		)
		if err != nil {
			log.Fatal(err)
		}
		arrProker = append(arrProker, proker)
	}
	return arrProker
}

func createProkerQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	nomorProgram := r.FormValue("noprog")
	namaProgram := r.FormValue("namaprog")
	suratKeterangan := r.FormValue("suket")

	db := connect()
	insert, err := db.Query("INSERT INTO proker VALUES(?,?,?)", nomorProgram, namaProgram, suratKeterangan)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer insert.Close()

	tpl.ExecuteTemplate(w, "return.html", nil)
}

func editProkerQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	nomorProgram := r.FormValue("noprog")
	namaProgram := r.FormValue("namaprog")
	suratKeterangan := r.FormValue("suket")

	db := connect()
	insert, err := db.Query("UPDATE proker SET namaProgram = ?, suratKeterangan = ? WHERE nomorProgram = ?", namaProgram, suratKeterangan, nomorProgram)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer insert.Close()

	tpl.ExecuteTemplate(w, "return.html", nil)
}

func deleteProkerQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	nomorProgram := r.FormValue("noprog")

	db := connect()
	insert, err := db.Query("DELETE FROM proker WHERE nomorProgram = ?", nomorProgram)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer insert.Close()

	tpl.ExecuteTemplate(w, "return.html", nil)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "admin" {
		tpl.ExecuteTemplate(w, "create", nil)
	} else {
		tpl.ExecuteTemplate(w, "wrong.html", nil)
	}
}
