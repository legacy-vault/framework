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

// server.go.

// HTTP Server.

package server

import (
	"fmt"
	"github.com/legacy-vault/framework/go/http_server/config"
	"github.com/legacy-vault/framework/go/http_server/exit"
	"github.com/legacy-vault/framework/go/http_server/model"
	"golang.org/x/net/context"
	"log"
	"net"
	"net/http"
	"time"
)

const StartUpErrorMonitorSleepTickMs = 500

const StartUpErrorMonitorMsgPrefix = "Starting the HTTP Server"
const StartUpErrorMonitorMsgTick = "."
const StartUpErrorMonitorMsgPostfixBad = "[FAILURE]"
const StartUpErrorMonitorMsgPostfixGood = "[DONE]"

const MsgStop = "HTTP Server Shutdown."

const BackgroundErrorMonitorMsgFormat = "HTTP Server Error: %v.\r\n"
const ReportDelayMs = 1000

const MsgBgErrorMonitorStart = "HTTP Server Background Error Monitor " +
	"has started."
const MsgBgErrorMonitorStop = "HTTP Server Background Error Monitor " +
	"has stopped."

const MsgFormatStartAddress = "HTTP Server has been started " +
	"at '%s' on '%s'."

type Server struct {
	HTTPServer                   http.Server
	StartUpErrorMonitoringPeriod time.Duration
	ErrorChan                    chan error
	AppQuitChannel               chan int
	ShutdownTimeout              time.Duration
	BTIH                         BTIHSettings
}

type BTIHSettings struct {
	TasksChannel       chan model.BTIHTask
	NewTasksAreAllowed *bool
}

type TimoutSetting struct {
	Idle       time.Duration
	Read       time.Duration
	ReadHeader time.Duration
	Write      time.Duration
	Shutdown   time.Duration
}

// Creates a new Server.
func New(
	host string,
	port string,
	timeoutSetting TimoutSetting,
	startUpErrorMonitoringPeriod int,
	appQuitChannel chan int,
	btihSettings BTIHSettings,
) *Server {

	var server *Server

	server = new(Server)
	server.initialize(
		host,
		port,
		timeoutSetting,
		startUpErrorMonitoringPeriod,
		appQuitChannel,
		btihSettings,
	)

	return server
}

// Prepares the Server for Work.
func (srv *Server) initialize(
	host string,
	port string,
	timeoutSetting TimoutSetting,
	startUpErrorMonitoringPeriod int,
	appQuitChannel chan int,
	btihSettings BTIHSettings,
) {

	// Set Address.
	srv.HTTPServer.Addr = net.JoinHostPort(host, port)

	// Set Timeouts.
	srv.HTTPServer.IdleTimeout = timeoutSetting.Idle
	srv.HTTPServer.ReadTimeout = timeoutSetting.Read
	srv.HTTPServer.ReadHeaderTimeout = timeoutSetting.ReadHeader
	srv.HTTPServer.WriteTimeout = timeoutSetting.Write

	// Set Router.
	srv.HTTPServer.Handler = http.HandlerFunc(srv.httpRouter)

	// Set Period for Monitoring of StartUp Errors.
	srv.StartUpErrorMonitoringPeriod = time.Second *
		time.Duration(startUpErrorMonitoringPeriod)

	// Prepare Error Channel.
	srv.ErrorChan = make(chan error, 1)

	// Save the Application Quit Channel.
	srv.AppQuitChannel = appQuitChannel

	// Shutdown Timeout.
	srv.ShutdownTimeout = timeoutSetting.Shutdown

	// BTIH Unit Tasks Channel.
	srv.BTIH.TasksChannel = btihSettings.TasksChannel
	srv.BTIH.NewTasksAreAllowed = btihSettings.NewTasksAreAllowed

	return
}

// Prepares HTTP Server Timeout Setting.
func TimeoutSetting(
	timeoutIdleSec uint,
	timeoutReadSec uint,
	timeoutReadHeaderSec uint,
	timeoutWriteSec uint,
	timeoutShutdownSec uint,
) TimoutSetting {

	var ts TimoutSetting

	ts.Idle = time.Second * time.Duration(timeoutIdleSec)
	ts.Read = time.Second * time.Duration(timeoutReadSec)
	ts.ReadHeader = time.Second * time.Duration(timeoutReadHeaderSec)
	ts.Write = time.Second * time.Duration(timeoutWriteSec)
	ts.Shutdown = time.Second * time.Duration(timeoutShutdownSec)

	return ts
}

// Starts the Server.
// Startup Period Error Monitoring is enabled.
// Background Error Monitoring is enabled.
func (srv *Server) Start() error {

	var err error
	var loop bool
	var msg string
	var sleepTick time.Duration
	var startTime time.Time
	var startTimeStr string
	var timeOfStart time.Time
	var timeOfStartUpMonitorEnd time.Time

	// Start the Server.
	timeOfStart = time.Now()
	go srv.asyncStart()

	// Monitor possible StartUp Errors.
	sleepTick = time.Millisecond *
		time.Duration(StartUpErrorMonitorSleepTickMs)
	timeOfStartUpMonitorEnd = timeOfStart.Add(srv.StartUpErrorMonitoringPeriod)

	// Report Prefix.
	fmt.Print(StartUpErrorMonitorMsgPrefix)

	// Monitor StartUp Errors.
	loop = true
	for loop {

		// Try to receive StartUp Error.
		select {

		case err = <-srv.ErrorChan:
			// Got an Error.

			// Report Prefix.
			fmt.Println(StartUpErrorMonitorMsgPostfixBad)
			time.Sleep(sleepTick) // Let the Print Happen!

			return err

		default:
			// No Error => Wait for Deadline.
			fmt.Print(StartUpErrorMonitorMsgTick)
			if time.Now().Before(timeOfStartUpMonitorEnd) {
				time.Sleep(sleepTick)
			} else {
				loop = false
			}
		}
	}

	// Start BackGround Error Monitor.
	go srv.bgErrorMonitor()

	// Report Postfix.
	fmt.Println(StartUpErrorMonitorMsgPostfixGood)

	// Log a delayed Address Report.
	startTime = time.Now()
	startTimeStr = startTime.Format(config.TimeFormat)
	if (config.App.Main.Verbose) {
		msg = fmt.Sprintf(
			MsgFormatStartAddress,
			srv.HTTPServer.Addr,
			startTimeStr,
		)
		go logDelayedReport(msg)
	}

	return nil
}

// Asynchronous Server Starter.
func (srv *Server) asyncStart() {

	var err error

	err = srv.HTTPServer.ListenAndServe()
	if err != nil {
		srv.ErrorChan <- err
	}
}

// Background HTTP Server Error Monitor.
func (srv *Server) bgErrorMonitor() {

	var err error
	var channelExists bool

	// Log a delayed Start Report.
	if (config.App.Main.Verbose) {
		go logDelayedReport(MsgBgErrorMonitorStart)
	}

	channelExists = true
	for channelExists {

		// Try to get an Error.
		err, channelExists = <-srv.ErrorChan
		if !channelExists {
			continue
		}

		// Report an Error.
		log.Printf(BackgroundErrorMonitorMsgFormat, err)

		// Skip the Error if it is not critical.
		if ErrorIsNotCritical(err) {
			continue
		}

		// Request the Application ShutDown.
		srv.AppQuitChannel <- exit.CodeHTTPServerFailure

		break
	}

	// Stop Report.
	log.Println(MsgBgErrorMonitorStop)
}

// Check whether an Error is not critical.
func ErrorIsNotCritical(err error) bool {

	if (err == http.ErrServerClosed) {
		return true
	}

	return false
}

// Logs a delayed Report.
func logDelayedReport(msg string) {

	var delay time.Duration

	delay = time.Millisecond * time.Duration(ReportDelayMs)
	time.Sleep(delay)
	log.Println(msg)
}

// Stops the Server.
func (srv *Server) Stop() error {

	var ctx context.Context
	var err error
	var timeoutShutdown time.Duration

	// Report.
	if (config.App.Main.Verbose) {
		log.Println(MsgStop)
	}

	timeoutShutdown = time.Second * srv.ShutdownTimeout
	ctx = context.Background()
	ctx, _ = context.WithTimeout(ctx, timeoutShutdown)
	err = srv.HTTPServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
