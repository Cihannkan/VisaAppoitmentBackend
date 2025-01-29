package services

import (
	"VisaAppoitmentBackend/models"
)

func MissionCountryList() []string {
	Appoitments := Getinfo()
	var missioncountrylist []string
	for _, appo := range Appoitments {
		if !contains(missioncountrylist, appo.Mission_country) {
			missioncountrylist = append(missioncountrylist, appo.Mission_country)
		}
	}
	return missioncountrylist
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func GetbyMissionCountry(ülkeadi string) []models.Appoitment {
	Appoitments := Getinfo()
	var bymissioncountry []models.Appoitment
	for _, appo := range Appoitments {
		if appo.Source_country == "Turkiye" && appo.Mission_country == ülkeadi {
			bymissioncountry = append(bymissioncountry, appo)
		}
	}
	return bymissioncountry
}
