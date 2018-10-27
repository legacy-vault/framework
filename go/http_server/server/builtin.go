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

// builtin.go.

// Application HTTP Server's Built-In Handlers.

package server

import (
	"bytes"
	"github.com/legacy-vault/framework/go/http_server/config"
	"github.com/legacy-vault/framework/go/http_server/stat"
	"net/http"
	"strconv"
	"time"
)

// HTTP Handler of URL='/system/appname'.
func (srv *Server) handlerAppName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(HeaderServer, config.HTTPServerName)
	w.WriteHeader(http.StatusOK)
	w.Write(replyAppNameBA)
	return
}

// HTTP Handler of URL='/system/ping'.
func (srv *Server) handlerPing(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(HeaderServer, config.HTTPServerName)
	w.WriteHeader(http.StatusOK)
	w.Write(replyPongBA)
	return
}

// HTTP Handler of URL='/system/ram'.
func (srv *Server) handlerRAMUsage(w http.ResponseWriter, r *http.Request) {

	var ramUsage uint64
	var ramUsageStr string

	// Prepare Data.
	ramUsage = stat.GetMemoryUsage()
	ramUsageStr = strconv.FormatUint(ramUsage, 10)

	w.Header().Set(HeaderServer, config.HTTPServerName)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ramUsageStr))
	return
}

// HTTP Handler of a non-existent URL.
func (srv *Server) handlerResourceNotFound(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set(HeaderServer, config.HTTPServerName)
	w.WriteHeader(http.StatusNotFound)
	w.Write(reply404BA)
	return
}

// HTTP Handler of a Root Page.
func (srv *Server) handlerRoot(w http.ResponseWriter, r *http.Request) {

	// Redirect to the '404' Page as we are not showing anything here.
	srv.handlerResourceNotFound(w, r)
	return
}

// HTTP Handler of URL='/system/statistics'.
func (srv *Server) handlerStatistics(w http.ResponseWriter, r *http.Request) {

	var buffer bytes.Buffer
	var ramUsageKiB uint64
	var ramUsageKiBStr string
	var replyBA []byte
	var reportTimeStr string
	var timeOfStart string
	var upTime int64
	var upTimeStr string
	var version string

	// Prepare Data.
	upTime = stat.GetTimeBeingAlive()
	timeOfStart = stat.StartTime.Format(config.TimeFormat)
	upTimeStr = strconv.FormatInt(upTime, 10)
	reportTimeStr = time.Now().Format(config.TimeFormat)
	version = config.AppVersion
	ramUsageKiB = stat.GetMemoryUsage() / 1024
	ramUsageKiBStr = strconv.FormatUint(ramUsageKiB, 10)

	// Compose a Report.
	buffer.WriteString("STATISTICS\r\n\r\n")
	buffer.WriteString("Service Name: " + config.AppName + ".\r\n")
	buffer.WriteString("Service Version: " + version + ".\r\n")
	buffer.WriteString("Time of Start: " + timeOfStart + ".\r\n")
	buffer.WriteString("Running Time (Seconds): " + upTimeStr + ".\r\n")
	buffer.WriteString("Operating System RAM Usage (KiB): " +
		ramUsageKiBStr + ".\r\n")
	buffer.WriteString("\r\n")
	buffer.WriteString("Report Time: " + reportTimeStr + ".\r\n")

	replyBA = buffer.Bytes()

	// Send the Reply.
	w.Header().Set(HeaderServer, config.HTTPServerName)
	w.WriteHeader(http.StatusOK)
	w.Write(replyBA)
	return
}

// HTTP Handler of URL='/system/uptime'.
func (srv *Server) handlerUptime(w http.ResponseWriter, r *http.Request) {

	var upTime int64
	var upTimeStr string

	// Prepare Data.
	upTime = stat.GetTimeBeingAlive()
	upTimeStr = strconv.FormatInt(upTime, 10)

	w.Header().Set(HeaderServer, config.HTTPServerName)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(upTimeStr))
	return
}

// HTTP Handler of URL='/system/version'.
func (srv *Server) handlerVersion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(HeaderServer, config.HTTPServerName)
	w.WriteHeader(http.StatusOK)
	w.Write(replyVersionBA)
	return
}
