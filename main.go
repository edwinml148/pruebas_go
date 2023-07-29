package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(uuid.New().String())
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola mundo"))
	})

	http.ListenAndServe(":3000", r)

}
