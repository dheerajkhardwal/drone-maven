// Copyright (c) 2020, the Drone Maven project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"testing"
	"github.com/drone-plugins/drone-plugin-lib/drone"
)

func TestPlugin(t *testing.T) {
	settings := Settings{
		UseCentral: false,
		Servers: []server{
			{ID: "releases", Username: "dheeraj@domain.io", Password: "HelloHello!"},
		},
		Repos: []repo{
			{ID: "releases", URL: "https://repo.domain.io/maven-releases", Releases: true, Snapshots: false},
		},
	}

	plugin := New(settings, drone.Pipeline{}, drone.Network{})
	plugin.Execute()
}
