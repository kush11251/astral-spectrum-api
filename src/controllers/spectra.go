package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"astral-spectrum-api/src/models"
)

func getSpectra(w http.ResponseWriter, r *http.Request) {
	spectra := []models.Spectrum{{ID: "1", Name: "Spectrum 1", Wavelength: 500.0}}
	json.NewEncoder(w).Encode(spectra)
}

func createSpectrum(w http.ResponseWriter, r *http.Request) {
	var spectrum models.Spectrum
	json.NewDecoder(r.Body).Decode(&spectrum)
	fmt.Println(spectrum)
}
