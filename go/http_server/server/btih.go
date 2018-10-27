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

// btih.go.

// Application HTTP Server's BTIH Handler.

package server

import (
	"fmt"
	"github.com/legacy-vault/framework/go/http_server/config"
	"github.com/legacy-vault/framework/go/http_server/helper"
	"github.com/legacy-vault/framework/go/http_server/model"
	"net/http"
)

const BTIHFileContentType = "application/x-bittorrent"

// HTTP Handler of URL='/<BTIHUrlPath>/.../'.
func (srv *Server) handlerBTIH(
	w http.ResponseWriter,
	r *http.Request,
	btih string,
) {

	var hdrContentDisposition string
	var fileExt string
	var fileName string
	var result model.BTIHTask
	var returnChannel chan model.BTIHTask
	var task model.BTIHTask

	// Prepare Data.
	fileExt = config.App.BTIHCache.FileExt

	// Prepare the BTIH Task.
	returnChannel = make(chan model.BTIHTask, 1)
	task = model.BTIHTask{
		BTIH:          btih,
		ReturnAddress: returnChannel,
	}

	// Send the Task to BTIH Queue if allowed.
	if *srv.BTIH.NewTasksAreAllowed != true {
		// Channel has been closed. No new Tasks are now allowed.
		w.Header().Set(HeaderServer, config.HTTPServerName)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}

	// Send the Task. The Task Manager will reply the Client.
	srv.BTIH.TasksChannel <- task

	// Receive the Results.
	result = <-returnChannel

	// Write the Reply to Client.
	if !result.FileExists {
		// File does not exist.
		w.Header().Set(HeaderServer, config.HTTPServerName)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}

	// File exists...

	// 1. HTTP Headers.
	fileName = result.BTIH + fileExt
	fileName = helper.PrepareStringForHTTPHeader(fileName)
	hdrContentDisposition = fmt.Sprintf(
		HeaderContentDispositionAttachment,
		fileName,
	)
	w.Header().Set(HeaderContentType, BTIHFileContentType)
	w.Header().Set(HeaderContentDisposition, hdrContentDisposition)
	w.Header().Set(HeaderServer, config.HTTPServerName)

	// 2. HTTP Status.
	w.WriteHeader(http.StatusOK)

	// 3. Content.
	w.Write(result.FileContents)

	return
}
