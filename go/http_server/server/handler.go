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

// handler.go.

// Application HTTP Server's custom Handlers.

package server

import (
	"github.com/legacy-vault/framework/go/http_server/config"
	"net/http"
)

// HTTP Handler of URL='...'.
func (srv *Server) handlerExample(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(HeaderServer, config.HTTPServerName)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte{})
	return
}
