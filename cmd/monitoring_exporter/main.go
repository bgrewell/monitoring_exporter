package main

import (
	"fmt"
	"github.com/bgrewell/monitoring_exporter/internal/collector"
	"github.com/bgrewell/monitoring_exporter/internal/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// Register all Prometheus metrics (SSH, System, etc.)
	collector.RegisterAllMetrics()

	// Start monitoring SSH logs
	go log.MonitorSSHLogs()

	// Additional system monitors or other monitoring components can be added here
	// go log.MonitorSystemLogs()

	// Start HTTP server to expose Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting monitoring exporter on :9090")
	http.ListenAndServe(":9090", nil)
}
