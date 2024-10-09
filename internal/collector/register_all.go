package collector

func RegisterAllMetrics() {
	// Register SSH metrics
	RegisterSSHMetrics()
	// Register Puppet metrics
	RegisterPuppetMetrics()
}
