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

// helper.go.

// Various helping Functions.

package helper

import "strings"

const DoubleQuote = `"`
const HeaderDoubleQuoteEscaped = `\"`
const Slash = "/"

// Prepares the String for a HTTP Headers.
func PrepareStringForHTTPHeader(s string) string {

	return strings.Replace(s, DoubleQuote, HeaderDoubleQuoteEscaped, -1)
}

// Turns a Path into Array of Path Components (Folders or Files).
func SplitPathIntoComponents(path string) []string {

	var components []string

	// Trim Slash at Corners.
	if strings.HasPrefix(path, Slash) {
		path = strings.TrimPrefix(path, Slash)
	}
	if strings.HasSuffix(path, Slash) {
		path = strings.TrimSuffix(path, Slash)
	}

	// Split by Slash.
	components = strings.Split(path, Slash)

	return components
}
