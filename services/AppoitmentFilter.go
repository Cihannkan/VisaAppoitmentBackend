package services

import (
	"VisaAppoitmentBackend/models"
	"encoding/json"
)

func MissionCountryList() string {
	Appoitments := Getinfo()
	var missioncountrylist []string
	for _, appo := range Appoitments {
		if !contains(missioncountrylist, appo.Mission_country) {
			missioncountrylist = append(missioncountrylist, appo.Mission_country)
		}
	}
	jsoncountrylist, _ := json.Marshal(missioncountrylist)
	return string(jsoncountrylist)
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func GetbyMissionCountry(ülkeadi string) string {
	Appoitments := Getinfo()
	var bysourcecountry []models.Appoitment
	for _, appo := range Appoitments {
		if appo.Source_country == "Turkiye" && appo.Mission_country == ülkeadi {
			bysourcecountry = append(bysourcecountry, appo)
		}
	}
	jsonbysourcecountry, _ := json.Marshal(bysourcecountry)
	return string(jsonbysourcecountry)
}
