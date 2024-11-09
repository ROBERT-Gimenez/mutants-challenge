package main

import (
	"Challenge/api/controller"
	"log"
	"net/http"
    "github.com/gorilla/mux"
)

func main() {

   router := mux.NewRouter()

   router.HandleFunc("/mutant", controller.CreateItem).Methods("POST")
   
    log.Println("Servidor iniciado en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}