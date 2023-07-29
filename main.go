package main

import (
	"fmt"
	"net/http"

	"github.com/edwinml148/pruebas_go/routes"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(uuid.New().String())
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)

}
