package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/spectra", getSpectra).Methods("GET")
	r.HandleFunc("/spectra", createSpectrum).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getSpectra(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Getting spectra...")
}

func createSpectrum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Creating spectrum...")
}
