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
// Creation Date:	2018-10-27.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// init.go.

// Application Initialization.

package main

import (
	"errors"
	"fmt"
	"github.com/legacy-vault/framework/go/http_server/btih"
	"github.com/legacy-vault/framework/go/http_server/config"
	"github.com/legacy-vault/framework/go/http_server/exit"
	"github.com/legacy-vault/framework/go/http_server/model"
	"github.com/legacy-vault/framework/go/http_server/server"
	"github.com/legacy-vault/framework/go/http_server/stat"
	"log"
	"os"
	"sync"
)

const ErrFormatInitializationNotRequired = "Initialization of '%s' was not" +
	" required but for some Reason it was started"

// Initializes the Application.
func initialize(app *Application) error {

	var err error

	// Initialize the Command Line Arguments.
	err = initCLA(app)
	if err != nil {
		log.Println(ErrCLAInit, err)
		return err
	}

	// Quit Infrastructure Preparations.
	err = initQuitInfrastructure(app)
	if err != nil {
		log.Println(ErrQuitInfrastructureInit, err)
		return err
	}

	// Initialize the Configuration File.
	err = initConfigFile()
	if err != nil {
		log.Println(ErrConfigFileInit, err)
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

	// Optional Functionality Units...

	// 1. BTIH.
	if (config.App.BTIHCache.IsEnabled) {
		err = initBTIH(app)
		if err != nil {
			log.Println(ErrBTIHInit, err)
			return err
		}
	}

	// Prepare the HTTP Server.
	err = initHTTPServer(
		config.App,
		app.HTTPServer,
		app,
	)
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

// Initializes the external Configuration File.
func initConfigFile() error {

	var cfg *config.XMLModelConfig
	var cfgFilePath string
	var err error

	// Check Configuration File Path.
	cfgFilePath = config.App.Main.ConfigurationFile
	if len(cfgFilePath) == 0 {
		// Configuration File is not available.

		// Switch off all optional Features (that are dependent on it).
		config.App.BTIHCache.IsEnabled = false

		// Quit normally.
		return nil
	}

	// Parse an external Configuration File.
	cfg, err = config.ParseExtConfigurationFile(cfgFilePath)
	if err != nil {
		return err
	}

	// Check BTIH Cache Root Path for Collisions.
	if cfg.BTIHCache.URLPath == server.PathSystem {
		err = errors.New(ErrBTIHCacheURLPathCollision)
		return err
	}

	// Save Configuration Data.
	config.App.BTIHCache.RootPath = cfg.BTIHCache.RootPath
	config.App.BTIHCache.URLPath = cfg.BTIHCache.URLPath
	config.App.BTIHCache.Capacity = cfg.BTIHCache.Capacity
	config.App.BTIHCache.TTL = cfg.BTIHCache.TTL
	config.App.BTIHCache.QueueSize = cfg.BTIHCache.QueueSize
	config.App.BTIHCache.FileExt = cfg.BTIHCache.FileExt
	config.App.BTIHCache.IsEnabled = true
	if config.App.Main.Verbose == true {
		fmt.Printf(
			MsgFormatFunctionalityEnabled,
			FuncUnitBTIHCache,
		)
	}

	return nil
}

// Initializes the HTTP Server.
func initHTTPServer(
	appCfg config.AppCfg,
	srv *server.Server,
	app *Application,
) error {

	var btihSettings server.BTIHSettings
	var timeoutSetting server.TimoutSetting
	var tmpSrv *server.Server

	timeoutSetting = server.TimeoutSetting(
		config.HTTPServerTimeoutIdle,
		config.HTTPServerTimeoutRead,
		config.HTTPServerTimeoutReadHeader,
		config.HTTPServerTimeoutWrite,
		config.HTTPServerTimeoutShutdown,
	)

	btihSettings = server.BTIHSettings{
		TasksChannel:       app.BTIH.Tasks,
		NewTasksAreAllowed: app.BTIH.NewTasksAreAllowed,
	}

	tmpSrv = server.New(
		appCfg.HTTP.Host,
		appCfg.HTTP.Port,
		timeoutSetting,
		config.HTTPServerStartupErrorMonitoringPeriod,
		app.QuitChannel,
		btihSettings,
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

// Starts the HTTP Server.
func initBTIH(app *Application) error {

	var err error

	// Is BTIH Enabled?
	if config.App.BTIHCache.IsEnabled == false {
		err = fmt.Errorf(
			ErrFormatInitializationNotRequired,
			FuncUnitBTIHCache,
		)
		return err
	}

	// 1.1. BTIH Cache Tasks Queue.
	app.BTIH.Tasks = make(
		chan model.BTIHTask,
		config.App.BTIHCache.QueueSize,
	)
	app.BTIH.NewTasksAreAllowed = new(bool)
	*app.BTIH.NewTasksAreAllowed = true

	// 1.2. Busy BTIH Manager.
	app.BTIH.BusyManagers = new(sync.WaitGroup)
	app.BTIH.BusyManagers.Add(1)

	// 1.3. BTIH Cache internals.
	err = btih.Init(
		config.App.BTIHCache.RootPath,
		config.App.BTIHCache.Capacity,
		config.App.BTIHCache.TTL,
		app.BTIH,
	)
	if err != nil {
		return err
	}

	return nil
}
