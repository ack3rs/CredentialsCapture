package config

import (
	"os"
	"time"
)

var THREADWAIT = 30 * time.Second
var HOSTNAME, _ = os.Hostname()
var VERSION = ""
var PROJECT = "CredialsCapture"

var configItems = map[string]string{
	"BIND":            "0.0.0.0:8080",
	"ENV":             "Development",
	"DatabaseDSN":     "sqluser:sqluser@tcp(127.0.0.1:3306)/credentials",
	"DatabaseDSN_LOG": "sqluser@Localhost", // This is the string written to the Logs, the FULL DSN contains the Password
}

// Get the Key.  First check the OS Env variables,  then the defaults used in development.
func Get(key string) string {

	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}

	if configItems[key] != "" {
		return configItems[key]
	}
	return ""
}
