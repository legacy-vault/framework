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

// BTIH Models.

package model

import (
	"sync"
)

type BTIHTask struct {
	BTIH          string
	FileExists    bool
	FileContents  []byte
	ReturnAddress chan BTIHTask
}

type BTIHData struct {
	// Channel for BTIH Unit Tasks.
	Tasks chan BTIHTask

	// BTIH busy Manager.
	BusyManagers *sync.WaitGroup

	// Are new Tasks allowed?
	NewTasksAreAllowed *bool
}
