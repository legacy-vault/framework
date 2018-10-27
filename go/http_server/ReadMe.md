# Application Framework with HTTP Server.


## Short Description.

This Application is a Framework with an extended HTTP Server Functionality.

## Full Description.

The Framework provides an Application with an HTTP Server.
The Framework has the following Features:
  -	Support for Command Line Arguments (used for Configuration);
  -	Support for an external Configuration File (used for additional Functionality Units);
  -	Application Initialization and Finalization are divided into Steps;
  -	Main Parameters are easily configurable using the Constants;
  -	Different Exit Codes for different Failures;
  -	HTTP Router with built-in Functionality, which includes:
    - Ping Handler (responds as "PONG" on '/ping' Request);
    - Application Name Handler;
    - Application Version Handler;
    - Application RAM Usage Handler;
    - Internal Application Statistics Handler;
    - 'Not Found' Page Handler;
    - 'Root Page' Handler;
  -	Configurable Timeouts, including Graceful Shutdown Timeout;
  -	Centralized Application Quit System which uses a Quit Event Channel;
  -	Operating System Signals Handling (SIGINT, SIGTERM);
  - Graceful Shutdown Functionality with configurable Timeout;
  -	HTTP Server Error Monitoring during Start-Up;
  -	HTTP Server Error Monitoring in Background.

The Framework provides an extended Functionality comparing with the Golang's 
built-in 'http' Package.

As an Example, this Framework has an additional Functionality Unit which uses 
the 'BTIH Cache' Library.

## Installation.

Import Commands:
```
go get -u "github.com/legacy-vault/framework/go/http_server"
```

## Usage.

The Framework is ready for Usage!<br />
Add your HTTP Handlers and enjoy the Server.
