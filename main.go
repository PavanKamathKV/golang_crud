package main

import (
	"log"
	"net/http"

	"golang_crud/routes"
	"golang_crud/services"
	"golang_crud/utility"
)

func main() {

	var db = utility.GetConnection()
	// defer db.Close()

	services.SetDB(db)
	var appRouter = routes.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))

}
