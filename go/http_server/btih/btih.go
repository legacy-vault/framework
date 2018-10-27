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

// BTIH File Cache.

package btih

import (
	"errors"
	"github.com/legacy-vault/framework/go/http_server/model"
	"github.com/legacy-vault/library/go/btih_cache"
)

const ErrStop = "BTIH Cache Stop Error"

var cache *btih_cache.BTIHCache

// Functionality Unit Initialization.
func Init(
	rootPath string,
	capacity uint64,
	ttl int64,
	btihData model.BTIHData,
) error {

	var err error

	// Initialize the Cache.
	cache, err = btih_cache.New(
		capacity,
		ttl,
		rootPath,
	)
	if err != nil {
		return err
	}

	// Start the Tasks Manager.
	go Manager(btihData)

	return nil
}

// Functionality Unit Finalization.
func Fin() error {

	var err error
	var ok bool

	// Stop the Cache.
	ok = cache.Stop()
	if !ok {
		err = errors.New(ErrStop)
		return err
	}

	return nil
}
