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

// config.go.

// Application Configuration.

package config

// Application's Global Settings.
const AppVersion = "1.0.0"
const AppName = "The Application Framework by McArcher"

// HTTP Server Settings.
const HTTPServerName = AppName + " " + AppVersion
const HTTPServerTimeoutIdle = 300                // Seconds.
const HTTPServerTimeoutRead = 300                // Seconds.
const HTTPServerTimeoutReadHeader = 300          // Seconds.
const HTTPServerTimeoutWrite = 300               // Seconds.
const HTTPServerStartupErrorMonitoringPeriod = 5 // Seconds.
const HTTPServerTimeoutShutdown = 60             // Seconds.

// Default Settings.
const DefaultConfigurationFile = ""
const DefaultHTTPServerHost = "0.0.0.0"
const DefaultHTTPServerPort = "2000"

// Formats.
const TimeFormat = "2006-01-02 15:04:05 MST"

type AppCfg struct {
	Main      MainCfg
	HTTP      HTTPCfg
	BTIHCache BTIHCacheCfg
}

type MainCfg struct {
	ConfigurationFile string
	Verbose           bool
}

type HTTPCfg struct {
	Host                string
	Port                string
	SystemStatIsEnabled bool
}

type BTIHCacheCfg struct {
	IsEnabled bool
	RootPath  string
	URLPath   string
	Capacity  uint64
	TTL       int64
	QueueSize uint64
	FileExt   string
}

var App AppCfg
