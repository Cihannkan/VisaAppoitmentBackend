package main

import (
	"VisaAppoitmentBackend/services"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/getmissioncountrylist", func(c fiber.Ctx) error {
		return c.JSON(services.MissionCountryList())
	})
	app.Get("/getappointmentbycountry/:country", func(c fiber.Ctx) error {
		country := c.Params("country")
		return c.JSON(services.GetbyMissionCountry(country))
	})

	app.Listen(":3000")
}
