//============================================================================//
//
// Copyright © 2018 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2018-10-24.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// init.go.

// Application Initialization.

package main

import (
	"github.com/legacy-vault/framework/go/http_server/config"
	"github.com/legacy-vault/framework/go/http_server/exit"
	"github.com/legacy-vault/framework/go/http_server/server"
	"github.com/legacy-vault/framework/go/http_server/stat"
	"log"
	"os"
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
