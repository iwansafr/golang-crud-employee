package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func UpdateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			name := r.FormValue("name")
			npwp := r.FormValue("npwp")
			address := r.FormValue("address")

			if name == "" || npwp == "" || address == "" {
				w.Write([]byte("Please fill all fields"))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err := db.Exec(
				"UPDATE employee SET name=?, npwp=?, address=? WHERE id=?",
				name, npwp, address, id,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusSeeOther)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")
			row := db.QueryRow("SELECT * FROM employee WHERE id = ?", id)
			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var employee Employee
			err := row.Scan(
				&employee.Id,
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["employee"] = employee

			tplErr := tmpl.Execute(w, data)
			if tplErr != nil {
				w.Write([]byte(tplErr.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
