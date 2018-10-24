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

// router.go.

// HTTP Server's Router.

package server

import (
	"github.com/legacy-vault/framework/go/http_server/config"
	"log"
	"net/http"
	"net/url"
)

const PathAppName = "/appname"
const PathPing = "/ping"
const PathStatistics = "/statistics"
const PathVersion = "/version"

var (
	reply404BA     = []byte("Resource is not found.")
	replyPongBA    = []byte("PONG")
	replyVersionBA = []byte(config.AppVersion)
	replyAppNameBA = []byte(config.AppName)
)

// Main HTTP Handler (Router).
func httpRouter(w http.ResponseWriter, r *http.Request) {

	var err error
	var reqURLPath string
	var reqURLFirstLetter byte
	var reqURLParams url.Values

	// Request's URL & Parameters.
	reqURLPath = r.URL.Path
	reqURLParams, err = url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		// Log.
		log.Println(err)
		return
	}

	// No Path is specified? => Root (Index) Page.
	if len(reqURLPath) <= 1 {
		handlerRoot(w, r)
		return
	}

	// Get the first Letter.
	reqURLFirstLetter = reqURLPath[1]

	// URL Parameters Processing.
	reqURLParams = reqURLParams //! <- This is a Plug.

	// Route a Request to a Method according to the URL requested.
	// Routing is done in Two Steps:
	// 1. We sort URLs by the first Letter, to speed up the Search;
	// 2. We search a Method according to the first Letter selected.
	switch reqURLFirstLetter {

	case 'a':
		if reqURLPath == PathAppName {
			handlerAppName(w, r)
			return
		} else {
			handlerResourceNotFound(w, r)
			return
		}

	case 'p':
		if reqURLPath == PathPing {
			handlerPing(w, r)
			return
		} else {
			handlerResourceNotFound(w, r)
			return
		}

	case 's':
		if reqURLPath == PathStatistics {
			handlerStatistics(w, r)
			return
		} else {
			handlerResourceNotFound(w, r)
			return
		}

	case 'v':
		if reqURLPath == PathVersion {
			handlerVersion(w, r)
			return
		} else {
			handlerResourceNotFound(w, r)
			return
		}

	default:
		handlerResourceNotFound(w, r)
		return
	}
}
