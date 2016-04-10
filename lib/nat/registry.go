// Copyright (C) 2015 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package nat

import (
	"time"
)

type DiscoverFunc func(renewal, timeout time.Duration) []Device

var providers []DiscoverFunc

func Register(provider DiscoverFunc) {
	providers = append(providers, provider)
}

func discoverAll(renewal, timeout time.Duration) map[string]Device {
	nats := make(map[string]Device)
	for _, discoverFunc := range providers {
		discoveredNATs := discoverFunc(renewal, timeout)
		for _, discoveredNAT := range discoveredNATs {
			nats[discoveredNAT.ID()] = discoveredNAT
		}
	}
	return nats
}
