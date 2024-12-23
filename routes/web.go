package routes

import (
	"database/sql"
	"net/http"

	"github.com/iwansafr/go_crud_employee/controllers"
)

func WebRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controllers.NewHelloWorldController())
	server.HandleFunc("/employee", controllers.IndexEmployeeController(db))
	server.HandleFunc("/employee/create", controllers.CreateEmployeeController(db))
	server.HandleFunc("/employee/update", controllers.UpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controllers.DeleteEmployeeController(db))
}
