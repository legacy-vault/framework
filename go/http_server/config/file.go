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

// file.go.

// Configuration File Handling.

package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

const ErrFormatFieldDataNotSet = "Configuration Field '%s' has no Data"

type XMLModelConfig struct {
	XMLName   xml.Name          `xml:"Configuration"` // Name of Root XML Tag.
	Label     string            `xml:"Label"`
	BTIHCache XMLModelBTIHCache `xml:"BTIHCache"`
}

type XMLModelBTIHCache struct {
	URLPath   string
	RootPath  string
	Capacity  uint64
	TTL       int64
	QueueSize uint64
	FileExt   string
}

// Parses an external Configuration File into an Object.
func ParseExtConfigurationFile(filePath string) (*XMLModelConfig, error) {

	var cfg *XMLModelConfig
	var err error
	var fileContents []byte

	// Read Contents of a Configuration File.
	fileContents, err = ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse Contents.
	cfg = new(XMLModelConfig)
	err = xml.Unmarshal(fileContents, cfg)
	if err != nil {
		return nil, err
	}

	// Check parsed Data.
	if len(cfg.Label) == 0 {
		err = fmt.Errorf(ErrFormatFieldDataNotSet, "Label")
		return nil, err
	}
	if len(cfg.BTIHCache.URLPath) == 0 {
		err = fmt.Errorf(ErrFormatFieldDataNotSet, "BTIHCache.URLPath")
		return nil, err
	}
	if len(cfg.BTIHCache.RootPath) == 0 {
		err = fmt.Errorf(ErrFormatFieldDataNotSet, "BTIHCache.RootPath")
		return nil, err
	}
	if cfg.BTIHCache.Capacity < 1 {
		err = fmt.Errorf(ErrFormatFieldDataNotSet, "BTIHCache.Capacity")
		return nil, err
	}
	if cfg.BTIHCache.TTL < 1 {
		err = fmt.Errorf(ErrFormatFieldDataNotSet, "BTIHCache.TTL")
		return nil, err
	}
	if cfg.BTIHCache.QueueSize < 1 {
		err = fmt.Errorf(ErrFormatFieldDataNotSet, "BTIHCache.QueueSize")
		return nil, err
	}

	return cfg, nil
}
