package router

import (
	"github.com/geekAshish/mongodb/controller"
	"github.com/gorilla/mux"
)



func Router() *mux.Router {
	router := mux.NewRouter();

	router.HandleFunc("/api/v1/movies", controller.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/v1/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/v1/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/v1/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/v1/movie", controller.DeleteAllMovie).Methods("DELETE")

	return router;
}