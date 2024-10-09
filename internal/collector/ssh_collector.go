package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	SshSessionStartTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ssh_session_start_total",
			Help: "Total number of SSH sessions started",
		},
		[]string{"user", "source_ip"},
	)
	SshSessionEndTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ssh_session_end_total",
			Help: "Total number of SSH sessions ended",
		},
		[]string{"user", "source_ip"},
	)
	SshSessionDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "ssh_session_duration_seconds",
			Help: "Duration of SSH sessions in seconds",
		},
		[]string{"user"},
	)
	SshActiveSessions = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssh_active_sessions",
			Help: "Number of active SSH sessions per user",
		},
		[]string{"user"},
	)
	SshFailedLoginTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ssh_failed_login_total",
			Help: "Total number of failed SSH login attempts",
		},
		[]string{"user", "source_ip"},
	)
)

func RegisterSSHMetrics() {
	// Register all Prometheus metrics
	prometheus.MustRegister(SshSessionStartTotal)
	prometheus.MustRegister(SshSessionEndTotal)
	prometheus.MustRegister(SshSessionDuration)
	prometheus.MustRegister(SshActiveSessions)
	prometheus.MustRegister(SshFailedLoginTotal)
}
