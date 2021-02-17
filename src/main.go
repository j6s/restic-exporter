package main

import (
	"flag"
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	config = flag.String("config", "restic-exporter.yaml", "Name of the config file to use")
	output = flag.String("output", "stats.txt", "File to export the stats to")
	resticBinary = flag.String("restic-bin", "restic", "Location of the restic binary to use (defaults to loading the one in your PATH)")
)

func collectMetrics(config Config) *prometheus.Registry {
	snapshot := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "restic_snapshot_timestamp",
	}, []string{"name"})
	registry := prometheus.NewRegistry()
	registry.Register(snapshot)

	for name, configItem := range config {
		restic := Restic{Binary: *resticBinary, Name: name, Repository: configItem.Repository, Password: configItem.Password}
		timestamp, err := restic.SnapshotTimestamp()
		if err != nil {
			log.Printf("[%s] <ERR> %s", name, err)
		}
		snapshot.WithLabelValues(name).Set(float64(timestamp))
	}

	return registry
}

func main() {
	flag.Parse()
	config := readConfig(*config)

	registry := collectMetrics(config)
	err := prometheus.WriteToTextfile(*output, registry)
	if err != nil {
		log.Fatal(err)
	}
}
