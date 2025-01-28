package main

import (
	"VisaAppoitmentBackend/services"
	"fmt"
)

func main() {
	x := services.Getinfo()
	fmt.Println(x)
}
