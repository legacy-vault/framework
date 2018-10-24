// init.go.

// Application Initialization.

package main

import (
	"log"
	"os"
	"tmp/ap/code/config" //!
	"tmp/ap/code/exit"   //!
	"tmp/ap/code/server" //!
	"tmp/ap/code/stat"   //!
)

// Initializes the Application.
func initialize(app *Application) error {

	var err error

	// Quit Infrastructure Preparations.
	err = initQuitInfrastructure(app)
	if err != nil {
		log.Println(ErrQuitInfrastructureInit, err)
		return err
	}

	// O.S. Signals Handler.
	err = initOSS(app)
	if err != nil {
		log.Println(ErrOSSHInit, err)
		return err
	}

	// Internal Statistics.
	err = stat.Init()
	if err != nil {
		log.Println(ErrStatInit, err)
		return err
	}

	// Prepare the HTTP Server.
	err = initHTTPServer(app.HTTPServer, app.QuitChannel)
	if err != nil {
		log.Println(ErrHTTPServerInit, err)
		return err
	}

	// Start the HTTP Server.
	err = startHTTPServer(app.HTTPServer)
	if err != nil {
		log.Println(ErrHTTPServerStart, err)
		return err
	}

	return nil
}

// Prepares the Quit Infrastructure.
func initQuitInfrastructure(app *Application) error {

	app.QuitChannel = make(chan int, 1)

	go appQuitMonitor(app)

	return nil
}

// Initializes the HTTP Server.
func initHTTPServer(srv *server.Server, appQuitChannel chan int) error {

	var timeoutSetting server.TimoutSetting
	var tmpSrv *server.Server

	timeoutSetting = server.TimeoutSetting(
		config.HTTPServerTimeoutIdle,
		config.HTTPServerTimeoutRead,
		config.HTTPServerTimeoutReadHeader,
		config.HTTPServerTimeoutWrite,
		config.HTTPServerTimeoutShutdown,
	)

	tmpSrv = server.New(
		config.HTTPServerHost,
		config.HTTPServerPort,
		timeoutSetting,
		config.HTTPServerStartupErrorMonitoringPeriod,
		appQuitChannel,
	)
	*srv = *tmpSrv

	return nil
}

// Starts the HTTP Server.
func startHTTPServer(srv *server.Server) error {

	var err error

	err = srv.Start()
	if err != nil {
		return err
	}

	return nil
}

// Monitors the Application Quit Channel.
func appQuitMonitor(app *Application) {

	var channelExists bool
	var err error
	var exitCode int

	channelExists = true
	for channelExists {

		// Try to get an Error.
		exitCode, channelExists = <-app.QuitChannel
		if !channelExists {
			continue
		}

		// Begin the Application ShutDown.
		err = Fin(app, exitCode)
		if err != nil {

			// Shutdown Failure? => Emergency Shutdown.
			log.Println(err)
			os.Exit(exit.CodeFinFailure)
		}

		break
	}
}
