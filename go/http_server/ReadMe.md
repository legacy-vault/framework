# Application Framework with HTTP Server.


## Short Description.

This Application is a Framework with an extended HTTP Server Functionality.

## Full Description.

The Framework provides an Application with an HTTP Server.
The Framework has the following Features:
	*	Application Initialization and Finalization are divided into Steps;
	*	Main Parameters are easily configurable using the Constants;
	*	Different Exit Codes for different Failures;
	*	HTTP Router with built-in Functionality, which includes:
		-	Ping Handler (responds as "PONG" on '/ping' Request;
		-	Application Name Handler;
		-	Application Version Handler;
		-	Internal Application Statistics Handler;
		- 	'Not Found' Page Handler;
		-	'Root Page' Handler;
	*	Configurable Timeouts, including Graceful Shutdown Timeout;
	*	Centralized Application Quit System which uses a Quit Event Channel;
	*	Operating System Signals Handling (SIGINT, SIGTERM);
	*	Graceful Shutdown Functionality with configurable Timeout;
	*	HTTP Server Error Monitoring during Start-Up;
	*	HTTP Server Error Monitoring in Background.

The Framework provides an extended Functionality comparing with the Golang's
built-in 'http' Package.

## Installation.

Import Commands:
```
go get "github.com/legacy-vault/framework/go/http_server"
```

## Usage.

The Framework is ready for Usage!
Add you HTTP Handlers and enjoy the Server.
