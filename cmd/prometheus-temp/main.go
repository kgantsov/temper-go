package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	temper "github.com/kgantsov/temper-go/pkg"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

type Metrics struct {
	temperature prometheus.Gauge
}

func NewMetrics(subsystem string) *Metrics {
	m := &Metrics{}

	m.temperature = prometheus.NewGauge(prometheus.GaugeOpts{
		Subsystem: subsystem,
		Name:      "temperature",
		Help:      "Current temperature",
	})
	prometheus.MustRegister(m.temperature)

	return m
}

func main() {
	flag.Parse()

	metrics := NewMetrics("TEMP")

	go func() {
		for {
			temp, err := temper.GetTemperature()

			if err == nil {
				metrics.temperature.Set(temp.Temperature)
			} else {
				log.Fatalf("Failed: %s", err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
