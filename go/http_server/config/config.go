// config.go.

// Application Configuration.

package config

// Application's Global Settings.
const AppVersion = "1.0.0"
const AppName = "The Application Framework by McArcher"

// HTTP Server Settings.
const HTTPServerHost = "0.0.0.0"
const HTTPServerPort = "2000"
const HTTPServerTimeoutIdle = 300                // Seconds.
const HTTPServerTimeoutRead = 300                // Seconds.
const HTTPServerTimeoutReadHeader = 300          // Seconds.
const HTTPServerTimeoutWrite = 300               // Seconds.
const HTTPServerStartupErrorMonitoringPeriod = 5 // Seconds.
const HTTPServerTimeoutShutdown = 60             // Seconds.
