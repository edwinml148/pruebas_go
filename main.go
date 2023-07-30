package main

import (
	"fmt"
	"net/http"

	"github.com/edwinml148/pruebas_go/db"
	"github.com/edwinml148/pruebas_go/models"
	"github.com/edwinml148/pruebas_go/routes"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(uuid.New().String())

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
