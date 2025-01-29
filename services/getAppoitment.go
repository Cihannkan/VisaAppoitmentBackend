package services

import (
	"VisaAppoitmentBackend/models"
	"encoding/json"
	"io"
	"net/http"
)

func Getinfo() []models.Appoitment {
	url := "https://api.schengenvisaappointments.com/api/visa-list/?format=json"
	response, _ := http.Get(url)
	responsebody, _ := io.ReadAll(response.Body)
	var Appoitments []models.Appoitment
	json.Unmarshal(responsebody, &Appoitments)
	return Appoitments
}
