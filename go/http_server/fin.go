// fin.go.

// Application Shutdown.

package main

import (
	"log"
	"os"
	"time"
	"tmp/ap/code/server" //!
	"tmp/ap/code/stat"   //!
)

const MsgShutdown = "Application Shutdown has started."
const MsgFormatStopReport = "Application has been stopped at %v.\r\n"

const ErrHTTPFin = "HTTP Server Stop Error:"
const ErrStatFin = "Internal Statistics Stop Error:"

// Finalizes the Application.
func Fin(app *Application, exitCode int) error {

	var err error
	var stopTime time.Time
	var stopTimeStr string

	log.Println(MsgShutdown)

	// Stop the HTTP Server.
	err = stopHTTPServer(app.HTTPServer)
	if err != nil {
		log.Println(ErrHTTPFin, err)
		return err
	}

	// Stop the internal Statistics.
	err = stat.Fin()
	if err != nil {
		log.Println(ErrStatFin, err)
		return err
	}

	// Report.
	stopTime = stat.StopTime
	stopTimeStr = stopTime.Format(stat.TimeFormat)
	log.Printf(MsgFormatStopReport, stopTimeStr)

	// Exit to Operating System.
	os.Exit(exitCode)
	return nil
}

// Stops the HTTP Server.
func stopHTTPServer(srv *server.Server) error {

	var err error

	err = srv.Stop()
	if err != nil {
		return err
	}

	return nil
}
