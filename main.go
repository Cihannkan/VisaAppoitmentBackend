package main

import (
	"VisaAppoitmentBackend/services"
	"fmt"
)

func main() {
	fmt.Println(services.GetbyMissionCountry("Norway"))
}
