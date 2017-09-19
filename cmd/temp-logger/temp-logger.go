package main

import (
	"github.com/kgantsov/temper-go"
	"log"
	"time"
)

func main() {
	for {
		temp, err := temper.GetTemperature()

		if err == nil {
			log.Printf("Temperature: %.2fK %.2fF %.2fC\n", temp+273.15, 9.0/5.0*temp+32, temp)
		} else {
			log.Fatalf("Failed: %s", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
