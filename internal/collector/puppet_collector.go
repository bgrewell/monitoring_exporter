package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	PuppetAgentState = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "puppet_agent_state",
			Help: "Current state of the Puppet agent: -2=not installed, -1=disabled, 0=unknown, 1=ok, 2=failed, 3=monitoring error.",
		},
		[]string{},
	)
	PuppetAgentLastRun = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "puppet_agent_last_run",
			Help: "Timestamp of the last Puppet agent run in Unix epoch format",
		},
		[]string{},
	)
	PuppetAgentResourcesTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "puppet_agent_resources_total",
			Help: "The total number of resources applied on the node",
		},
		[]string{},
	)
	PuppetAgentResourcesChanged = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "puppet_agent_resources_changed",
			Help: "The number of resources that changed on the last run",
		},
		[]string{},
	)
	PuppetAgentApplyTotalTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "puppet_agent_apply_total_time",
			Help: "Total time taken to apply the last Puppet agent run in seconds",
		},
		[]string{},
	)
)

func RegisterPuppetMetrics() {
	// Register all Prometheus metrics
	prometheus.MustRegister(PuppetAgentState)
	prometheus.MustRegister(PuppetAgentLastRun)
	prometheus.MustRegister(PuppetAgentResourcesTotal)
	prometheus.MustRegister(PuppetAgentResourcesChanged)
	prometheus.MustRegister(PuppetAgentApplyTotalTime)
}
