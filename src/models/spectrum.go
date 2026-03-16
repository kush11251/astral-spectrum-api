package models

import (
	"encoding/json"
)

type Spectrum struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Wavelength float64 `json:"wavelength"`
}
