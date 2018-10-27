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

// main.go.

// Application's Entry Point.

package main

import (
	"github.com/legacy-vault/framework/go/http_server/exit"
	"github.com/legacy-vault/framework/go/http_server/server"
	"log"
	"os"
)

// Error Messages.
const (
	ErrError                     = "Error:"
	ErrInit                      = "Initialization Error:"
	ErrStart                     = "Start Error:"
	ErrHTTPServerInit            = "HTTP Server " + ErrInit
	ErrHTTPServerStart           = "HTTP Server " + ErrStart
	ErrOSSHInit                  = "O.S. Signals Handler " + ErrInit
	ErrQuitInfrastructureInit    = "Application Quit Infrastructure " + ErrInit
	ErrStatInit                  = "Internal Statistics " + ErrInit
	ErrCLAInit                   = "Command Line Arguments " + ErrInit
	ErrConfigFileInit            = "Configuration File " + ErrInit
	ErrBTIHCacheURLPathCollision = "BTIH Cache URL Path Collision Error"
	ErrBTIHInit                  = "BTIH Cache " + ErrInit
)

// Normal Messages and their Formats.
const (
	MsgFormatFunctionalityEnabled = "<%s> Functionality has been Enabled.\r\n"
	MsgFormatUnitStopped          = "<%s> Unit has been Stopped.\r\n"
)

// Program's Entry Point.
func main() {

	var app Application
	var err error
	var void chan bool

	app.HTTPServer = new(server.Server)
	err = initialize(&app)
	if err != nil {
		log.Println(ErrInit, err)
		os.Exit(exit.CodeInitFailure)
	}

	// Wait forever.
	void = make(chan bool, 1)
	_ = <-void
	os.Exit(exit.CodeOK)
}
