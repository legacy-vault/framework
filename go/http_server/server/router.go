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

// router.go.

// HTTP Server's Router.

package server

import (
	"github.com/legacy-vault/framework/go/http_server/config"
	"github.com/legacy-vault/framework/go/http_server/helper"
	"log"
	"net/http"
)

const PathSystem = "system"
const PathSystemAppName = "appname"
const PathSystemPing = "ping"
const PathSystemRAMUsage = "ram"
const PathSystemStatistics = "statistics"
const PathSystemUptime = "uptime"
const PathSystemVersion = "version"

var (
	reply404BA     = []byte("Resource is not found.")
	replyPongBA    = []byte("PONG")
	replyVersionBA = []byte(config.AppVersion)
	replyAppNameBA = []byte(config.AppName)
)

// Main HTTP Handler (Router).
func httpRouter(w http.ResponseWriter, r *http.Request) {

	var pathComponent1 string
	var pathComponent2 string
	var pathComponents []string
	var reqURLPath string

	// Request's URL & Parameters.
	reqURLPath = r.URL.Path

	// No Path is specified? => Root (Index) Page.
	if len(reqURLPath) <= 1 {
		handlerRoot(w, r)
		return
	}

	// Split Path into Components.
	pathComponents = helper.SplitPathIntoComponents(reqURLPath)

	if len(pathComponents) >= 1 {
		pathComponent1 = pathComponents[0]
	} else {
		handlerRoot(w, r)
		return
	}

	if (config.App.HTTP.SystemStatIsEnabled) && (pathComponent1 == PathSystem) {

		// System Handlers.

		// Bad Request?
		if len(pathComponents) < 2 {
			handlerResourceNotFound(w, r)
			return
		}

		pathComponent2 = pathComponents[1]
		switch pathComponent2 {

		case PathSystemPing:
			handlerPing(w, r)
			return

		case PathSystemUptime:
			handlerUptime(w, r)
			return

		case PathSystemRAMUsage:
			handlerRAMUsage(w, r)
			return

		case PathSystemAppName:
			handlerAppName(w, r)
			return

		case PathSystemVersion:
			handlerVersion(w, r)
			return

		case PathSystemStatistics:
			handlerStatistics(w, r)
			return

		default:
			handlerResourceNotFound(w, r)
			return
		}
	}

	// Normal Path.
	if (config.App.Main.Verbose) {
		log.Println(pathComponent1)
	}

	handlerResourceNotFound(w, r)
	return
}
