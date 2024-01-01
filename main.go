package main

import (
	"fmt"
	"log"
	"net/http"
	"udon-web/controllers"

	"github.com/gorilla/mux"
)

func main() {
	LoadAppConfig()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	RegisterUdonRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterUdonRoutes(router *mux.Router) {
	var muxBase = "/api/udons"
	router.HandleFunc(muxBase, controllers.GetUdons).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.GetUdonById).Methods("GET")
	router.HandleFunc(muxBase, controllers.CreateUdon).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.UpdateUdon).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.DeleteUdon).Methods("DELETE")
}