package controller

import (
	"Challenge/api/models"
	"Challenge/api/service"
	"encoding/json"
	"net/http"
	"sync"
	)

var (
	mu     sync.Mutex
	mutantService models.MutantService = &service.MutantService{}
	)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var mutant models.Mutants
	if err := json.NewDecoder(r.Body).Decode(&mutant); err != nil {
		http.Error(w, "Error al procesar la solicitud", http.StatusBadRequest)
		return
	}
	isMutant := mutantService.IsMutant(mutant.Adn)

    mu.Lock()
	mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(isMutant)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	

    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("hola")
}