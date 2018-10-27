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

// manager.go.

// BTIH Cache Tasks Queue Manager.

package btih

import (
	"github.com/legacy-vault/framework/go/http_server/model"
	"log"
)

const MsgManagerStop = "BTIH Task Manager has stopped."

// Processes incoming Tasks.
func Manager(btihData model.BTIHData) {

	var err error
	var fileContents []byte
	var queue chan model.BTIHTask
	var reply model.BTIHTask
	var task model.BTIHTask

	// Prepare Reply.
	reply.ReturnAddress = nil

	// 'Work is done' Report at Exit.
	defer btihData.BusyManagers.Done()

	// Get Tasks while they are available.
	queue = btihData.Tasks
	for task = range queue {

		// Process an incoming Task.

		// Prepare Reply's Data.
		reply.BTIH = task.BTIH

		// Get File Contents from BTIH Cache.
		fileContents, err = cache.GetFileByBTIH(task.BTIH)
		if err != nil {

			// File is not accessible.

			// Set the Reply.
			reply.FileExists = false
			reply.FileContents = []byte{}

			// Send the Result back to Sender.
			task.ReturnAddress <- reply

			continue
		}

		// File is accessible.

		// Set the Reply.
		reply.FileExists = true
		reply.FileContents = fileContents

		// Send the Result back to Sender.
		task.ReturnAddress <- reply

		continue
	}

	// Exit.
	log.Println(MsgManagerStop)
}
