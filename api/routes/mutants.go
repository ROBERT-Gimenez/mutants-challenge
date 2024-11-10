package routes

import(
	"Challenge/api/controller"
	"Challenge/api/repository"
	"Challenge/api/service"
	"github.com/gorilla/mux"
)

func MutantsRoutes(router *mux.Router) {

	mutantRepo , _ := repository.NewMutantRepository()
	mutantServices := service.NewMutantService(mutantRepo)
	mutantController := controller.NewMutantController(mutantServices)

	router.HandleFunc("/mutant", mutantController.PostMutantDNA).Methods("POST")
	router.HandleFunc("/stats", mutantController.GetStats).Methods("GET")

}
