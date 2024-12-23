package main

import (
	"fmt"
	"net/http"

	"github.com/iwansafr/go_crud_employee/database"
	"github.com/iwansafr/go_crud_employee/routes"
)

func main() {
	db := database.InitDatabase()
	server := http.NewServeMux()

	routes.WebRoutes(server, db)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", server)
}
