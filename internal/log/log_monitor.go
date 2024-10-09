package log

import (
	"bufio"
	"fmt"
	"github.com/bgrewell/monitoring_exporter/internal/collector"
	"os"
	"regexp"
	"runtime"
	"time"
)

// Regular expressions for log parsing
var (
	acceptedAuthRegex     = regexp.MustCompile(`(?mi)Accepted (\w+) for (\w+) from ([\d\.]+)`)
	connectionClosedRegex = regexp.MustCompile(`(?mi)Disconnected from user (\w+) ([\d\.]+)`)
	sessionOpenedRegex    = regexp.MustCompile(`(?mi)session opened for user (\w+)`)
	sessionClosedRegex    = regexp.MustCompile(`(?mi)session closed for user (\w+)`)
)

func getLogFilePath() string {
	switch runtime.GOOS {
	case "linux":
		if _, err := os.Stat("/var/log/auth.log"); err == nil {
			return "/var/log/auth.log"
		}
		if _, err := os.Stat("/var/log/secure"); err == nil {
			return "/var/log/secure"
		}
	default:
		fmt.Println("Unsupported OS.")
		os.Exit(1)
	}
	return ""
}

func MonitorSSHLogs() {
	filePath := getLogFilePath()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		parseSSHLogLine(line)
	}
}

// Parse log line and update Prometheus metrics
func parseSSHLogLine(line string) {
	if acceptedAuthRegex.MatchString(line) {
		// Handle successful login (password or public key)
		authType, user, sourceIP := extractAuthInfo(line)
		handleSuccessfulLogin(user, sourceIP, authType)
	} else if connectionClosedRegex.MatchString(line) {
		// Handle closed connection
		user, sourceIP := extractClosedConnectionInfo(line)
		handleConnectionClosed(user, sourceIP)
		// TODO: Uncomment this block to handle failed logins
		//// Handle failed login (key or password)
		//user, sourceIP := extractFailedLoginInfo(line)
		//handleFailedLogin(user, sourceIP)
	} else if sessionOpenedRegex.MatchString(line) {
		// Handle session start
		user := extractSessionUser(line, sessionOpenedRegex)
		handleSessionOpened(user)
	} else if sessionClosedRegex.MatchString(line) {
		// Handle session closure
		user := extractSessionUser(line, sessionClosedRegex)
		handleSessionClosed(user)
	}
}

// Extract auth type, user, and IP for successful login
func extractAuthInfo(line string) (string, string, string) {
	matches := acceptedAuthRegex.FindStringSubmatch(line)
	if len(matches) > 3 {
		return matches[1], matches[2], matches[3] // authType, user, sourceIP
	}
	return "", "", ""
}

func extractClosedConnectionInfo(line string) (string, string) {
	matches := connectionClosedRegex.FindStringSubmatch(line)
	if len(matches) > 2 {
		return matches[1], matches[2] // user, sourceIP
	}
	return "", ""
}

// Extract user and IP for failed login
func extractFailedLoginInfo(line string) (string, string) {
	matches := connectionClosedRegex.FindStringSubmatch(line)
	if len(matches) > 2 {
		return matches[1], matches[2] // user, sourceIP
	}
	return "", ""
}

// Extract user for session events (opened/closed)
func extractSessionUser(line string, regex *regexp.Regexp) string {
	matches := regex.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1] // user
	}
	return ""
}

// Handle a successful login event
func handleSuccessfulLogin(user, sourceIP, authType string) {
	// Increment session start total for Prometheus metrics
	collector.SshSessionStartTotal.WithLabelValues(user, sourceIP).Inc()

	// Increment active sessions for the user
	collector.SshActiveSessions.WithLabelValues(user).Inc()
}

// Handle connection closed
func handleConnectionClosed(user, sourceIP string) {
	// Increment session end total for Prometheus metrics
	collector.SshSessionEndTotal.WithLabelValues(user, sourceIP).Inc()

	// Decrement active sessions for the user
	collector.SshActiveSessions.WithLabelValues(user).Dec()
}

// Handle a failed login event
func handleFailedLogin(user, sourceIP string) {
	// Increment failed login counter (you could add a failed login metric if needed)
	collector.SshFailedLoginTotal.WithLabelValues(user, sourceIP).Inc()
}

// Handle session opened event
func handleSessionOpened(user string) {
	// Increment active sessions'
	//collector.SshActiveSessions.WithLabelValues(user).Inc()
}

// Handle session closed event
func handleSessionClosed(user string) {
	// Decrement active sessions'
	//collector.SshActiveSessions.WithLabelValues(user).Dec()
}
