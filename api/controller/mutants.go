package controller

import (
	"Challenge/api/models"
	"Challenge/api/service"
	"encoding/json"
	"log"
	"net/http"
)

type MutantController struct {
	mutantServices *service.MutantService
}

func NewMutantController(service *service.MutantService) *MutantController {
	return &MutantController{mutantServices: service}
}


func (m *MutantController) PostMutantDNA(w http.ResponseWriter, r *http.Request) {
	var mutant models.Mutants

	if err := json.NewDecoder(r.Body).Decode(&mutant); err != nil {
		http.Error(w, "Error al procesar la solicitud", http.StatusBadRequest)
		return
	}

	isMutant, err := m.mutantServices.IsMutant(mutant.Adn)
	if err != nil {
		http.Error(w, "Error interno en el servicio", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if isMutant {
		w.WriteHeader(http.StatusOK)
	}else{
		w.WriteHeader(http.StatusForbidden)
	}
	err = json.NewEncoder(w).Encode(isMutant)
	if err != nil {
		log.Printf("Error al codificar la respuesta: %v", err)
	}
}

func (m *MutantController) GetStats(w http.ResponseWriter, r *http.Request) {
	stats , err := m.mutantServices.GetStatsMutant()
	if err != nil {
		log.Printf("Error al obtener las stats: %v", err)
	}
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(stats)
}