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

// app.go.

// Application.

package main

import (
	"github.com/legacy-vault/framework/go/http_server/model"
	"github.com/legacy-vault/framework/go/http_server/server"
	"os"
)

type Application struct {
	// HTTP Server.
	HTTPServer *server.Server

	// Channel accepting Application Quit Signals.
	QuitChannel chan int

	// Channel accepting O.S. Signals for Service Termination.
	OSTermSignalsChannel chan os.Signal

	// BTIH Data.
	BTIH model.BTIHData
}
