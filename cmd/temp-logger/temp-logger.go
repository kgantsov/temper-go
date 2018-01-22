package main

import (
	"log"
	"time"

	"github.com/kgantsov/temper-go/pkg"
)

func main() {
	for {
		temp, err := temper.GetTemperature()

		if err == nil {
			log.Printf("Temperature: %.2f°K %.2f°F %.2f°C\n", temp+273.15, 9.0/5.0*temp+32, temp)
		} else {
			log.Fatalf("Failed: %s", err)
		}
		time.Sleep(1 * time.Second)
	}
}
