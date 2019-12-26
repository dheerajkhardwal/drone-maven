// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import "testing"

import "github.com/drone-plugins/drone-plugin-lib/pkg/plugin"
import "github.com/drone-plugins/drone-plugin-lib/pkg/urfave"

func TestPlugin(t *testing.T) {
	settings := Settings{
		UseCentral: true,
		Servers: []server{
			{ID: "Nexus", Username: "dheeraj@domain.io", Password: "HelloHello!"},
		},
		Repos: []repo{
			{ID: "Nexus", URL: "https://repo.domain.io/repo/maven", Releases: true, Snapshots: true},
		},
	}

	New(settings, plugin.Pipeline{}, urfave.Network{}).Exec()
}
