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

// fin.go.

// Application Shutdown.

package main

import (
	"github.com/legacy-vault/framework/go/http_server/config"
	"github.com/legacy-vault/framework/go/http_server/server"
	"github.com/legacy-vault/framework/go/http_server/stat"
	"log"
	"os"
	"time"
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

	if (config.App.Main.Verbose) {
		log.Println(MsgShutdown)
	}

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
	if (config.App.Main.Verbose) {
		stopTime = stat.StopTime
		stopTimeStr = stopTime.Format(stat.TimeFormat)
		log.Printf(MsgFormatStopReport, stopTimeStr)
	}

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
