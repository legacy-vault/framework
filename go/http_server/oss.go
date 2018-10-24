// oss.go.

// Operating System Signals Handler.

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tmp/ap/code/exit" //!
)

const MsgSIGTERM = "'SIGTERM' Signal is received."
const MsgSIGINT = "'SIGINT' Signal is received."

const OSSignalsTermChanSize = 1

// Initializes O.S. Signals Handling.
func initOSS(app *Application) error {

	app.OSTermSignalsChannel = make(chan os.Signal, OSSignalsTermChanSize)
	signal.Notify(app.OSTermSignalsChannel, syscall.SIGINT, syscall.SIGTERM)

	go handlerOfTermination(
		app.OSTermSignalsChannel,
		app.QuitChannel,
	)

	return nil
}

// Handler of terminating Signals: 'SIGTERM' and 'SIGINT'.
func handlerOfTermination(
	signals chan os.Signal,
	appQuitChannel chan int,
) {

	var sig os.Signal

	for sig = range signals {

		switch sig {

		case syscall.SIGTERM:

			// Log.
			log.Println(MsgSIGTERM)

			// Request the Application Shutdown.
			appQuitChannel <- exit.CodeSignalSIGTERM

		case syscall.SIGINT:

			// Log.
			log.Println(MsgSIGINT)

			// Request the Application Shutdown.
			appQuitChannel <- exit.CodeSignalSIGINT
		}
	}
}
