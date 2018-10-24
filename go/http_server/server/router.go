// router.go.

// HTTP Server's Router.

package server

import (
	"log"
	"net/http"
	"net/url"
	"tmp/ap/code/config" //!
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
