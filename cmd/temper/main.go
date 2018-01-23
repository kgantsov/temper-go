package main

import (
	"log"

	"github.com/kgantsov/temper-go/pkg"
)

func main() {
	temp, err := temper.GetTemperature()

	if err == nil {
		log.Printf(
			"Temperature: %.2f°K %.2f°F %.2f°C\n",
			temp.Temperature+273.15,
			9.0/5.0*temp.Temperature+32,
			temp.Temperature,
		)
	} else {
		log.Fatalf("Failed: %s", err)
	}
}
