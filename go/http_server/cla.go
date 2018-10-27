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

// cla.go.

// Command Line Arguments Handling.

package main

import (
	"flag"
	"github.com/legacy-vault/framework/go/http_server/config"
)

// Names of Command Line Arguments (Keys).
const CLANameConfigurationFile = "config-file"
const CLANameHTTPHost = "host"
const CLANameHTTPPort = "port"
const CLANameSystemStatistics = "ss"
const CLANameVerbose = "v"

// Default values of Command Line Arguments.
const CLADefaultValueConfigurationFile = config.DefaultConfigurationFile
const CLADefaultValueHTTPHost = config.DefaultHTTPServerHost
const CLADefaultValueHTTPPort = config.DefaultHTTPServerPort
const CLADefaultValueSystemStatistics = true
const CLADefaultValueVerbose = true

// Hint Texts of Command Line Arguments.
const CLAHintConfigurationFile = "External Configuration File"
const CLAHintHTTPHost = "HTTP Server's Host Name or IP Address"
const CLAHintHTTPPort = "HTTP Server's Port"
const CLAHintSystemStatistics = "System Statistics"
const CLAHintVerbose = "Verbose Mode"

// Initializes the Command Line Arguments.
func initCLA(app *Application) error {

	var claConfigFile *string
	var claHTTPHost *string
	var claHTTPPort *string
	var claSystemStats *bool
	var claVerbose *bool

	// Prepare C.L.A. Parameters.
	claConfigFile = flag.String(
		CLANameConfigurationFile,
		CLADefaultValueConfigurationFile,
		CLAHintConfigurationFile,
	)
	claHTTPHost = flag.String(
		CLANameHTTPHost,
		CLADefaultValueHTTPHost,
		CLAHintHTTPHost,
	)
	claHTTPPort = flag.String(
		CLANameHTTPPort,
		CLADefaultValueHTTPPort,
		CLAHintHTTPPort,
	)
	claSystemStats = flag.Bool(
		CLANameSystemStatistics,
		CLADefaultValueSystemStatistics,
		CLAHintSystemStatistics,
	)
	claVerbose = flag.Bool(
		CLANameVerbose,
		CLADefaultValueVerbose,
		CLAHintVerbose,
	)

	// Read Flags.
	flag.Parse()
	config.App.Main.ConfigurationFile = *claConfigFile
	config.App.HTTP.Host = *claHTTPHost
	config.App.HTTP.Port = *claHTTPPort
	config.App.HTTP.SystemStatIsEnabled = *claSystemStats
	config.App.Main.Verbose = *claVerbose

	return nil
}
