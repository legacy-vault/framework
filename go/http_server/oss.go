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

// oss.go.

// Operating System Signals Handler.

package main

import (
	"github.com/legacy-vault/framework/go/http_server/exit"
	"log"
	"os"
	"os/signal"
	"syscall"
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
