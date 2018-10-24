// main.go.

// Application's Entry Point.

package main

import (
	"log"
	"os"
	"tmp/ap/code/exit"   //!
	"tmp/ap/code/server" //!
)

const ErrInit = "Initialization Error:"
const ErrHTTPServerInit = "HTTP Server Initialization Error:"
const ErrHTTPServerStart = "HTTP Server Start Error:"
const ErrOSSHInit = "O.S. Signals Handler Initialization Error:"
const ErrQuitInfrastructureInit = "Application Quit Infrastructure Error:"
const ErrStatInit = "Internal Statistics Initialization Error:"

// Program's Entry Point.
func main() {

	var app Application
	var err error
	var void chan bool

	app.HTTPServer = new(server.Server)
	err = initialize(&app)
	if err != nil {
		log.Println(ErrInit, err)
		os.Exit(exit.CodeInitFailure)
	}

	// Wait forever.
	void = make(chan bool, 1)
	_ = <-void
	os.Exit(exit.CodeOK)
}
