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

// stat.go.

// Application's Statistics.

package stat

import (
	"runtime"
	"time"
)

var StartTime time.Time
var StartTimestamp int64

var StopTime time.Time
var StopTimestamp int64

// Initializes Statistics.
func Init() error {

	StartTime = time.Now()
	StartTimestamp = StartTime.Unix()

	return nil
}

// Finalizes Statistics.
func Fin() error {

	StopTime = time.Now()
	StopTimestamp = StopTime.Unix()

	return nil
}

// Returns the Duration (in Seconds) of the Service being alive ("up-time").
func GetTimeBeingAlive() int64 {

	var tsNow int64
	var upTime int64

	tsNow = time.Now().Unix()
	upTime = tsNow - StartTimestamp

	return upTime
}

// Returns Application's Memory Usage Statistics.
func GetMemoryUsage() uint64 {

	var m runtime.MemStats
	var ramUsedfromOS uint64

	runtime.ReadMemStats(&m)
	ramUsedfromOS = m.Sys

	return ramUsedfromOS
}
