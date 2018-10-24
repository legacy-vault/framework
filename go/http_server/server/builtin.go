// builtin.go.

// Application HTTP Server's Built-In Handlers.

package server

import (
	"bytes"
	"net/http"
	"strconv"
	"time"
	"tmp/ap/code/config" //!
	"tmp/ap/code/stat"   //!
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
