// app.go.

// Application.

package main

import (
	"os"
	"tmp/ap/code/server" //!
)

type Application struct {
	// HTTP Server.
	HTTPServer *server.Server

	// Channel accepting Application Quit Signals.
	QuitChannel chan int

	// Channel accepting O.S. Signals for Service Termination.
	OSTermSignalsChannel chan os.Signal
}
