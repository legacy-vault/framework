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

// HTTP Handler of a Root Page.
func handlerRoot(w http.ResponseWriter, r *http.Request) {

	// Redirect to the '404' Page as we are not showing anything here.
	handlerResourceNotFound(w, r)
	return
}

// HTTP Handler of a non-existent URL.
func handlerResourceNotFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	w.Write(reply404BA)
	return
}

// HTTP Handler of URL='/ping'.
func handlerPing(w http.ResponseWriter, r *http.Request) {

	w.Write(replyPongBA)
	return
}

// HTTP Handler of URL='/statistics'.
func handlerStatistics(w http.ResponseWriter, r *http.Request) {

	var buffer bytes.Buffer
	var replyBA []byte
	var reportTimeStr string
	var timeOfStart string
	var upTime int64
	var upTimeStr string
	var version string

	// Prepare Data.
	upTime = stat.GetTimeBeingAlive()
	timeOfStart = stat.StartTime.Format(stat.TimeFormat)
	upTimeStr = strconv.FormatInt(upTime, 10)
	reportTimeStr = time.Now().Format(stat.TimeFormat)
	version = config.AppVersion

	// Compose a Report.
	buffer.WriteString("STATISTICS\r\n\r\n")
	buffer.WriteString("Service Name: " + config.AppName + ".\r\n")
	buffer.WriteString("Service Version: " + version + ".\r\n")
	buffer.WriteString("Time of Start: " + timeOfStart + ".\r\n")
	buffer.WriteString("Running Time (Seconds): " + upTimeStr + ".\r\n")
	buffer.WriteString("\r\n")
	buffer.WriteString("Report Time: " + reportTimeStr + ".\r\n")

	replyBA = buffer.Bytes()

	// Send the Reply.
	w.Write(replyBA)
	return
}

// HTTP Handler of URL='/version'.
func handlerVersion(w http.ResponseWriter, r *http.Request) {

	w.Write(replyVersionBA)
	return
}

// HTTP Handler of URL='/appname'.
func handlerAppName(w http.ResponseWriter, r *http.Request) {

	w.Write(replyAppNameBA)
	return
}
