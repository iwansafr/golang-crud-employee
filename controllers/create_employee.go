package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.FormValue("name")
			npwp := r.FormValue("npwp")
			address := r.FormValue("address")

			if name == "" || npwp == "" || address == "" {
				w.Write([]byte("Please fill all fields"))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err := db.Exec(
				"INSERT INTO employee (name, npwp, address) VALUES (?,?,?)",
				name, npwp, address,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusSeeOther)
			return
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tplErr := tmpl.Execute(w, nil)
			if tplErr != nil {
				w.Write([]byte(tplErr.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
