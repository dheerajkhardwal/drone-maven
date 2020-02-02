// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"os/user"
	"path"
	"text/template"
	"fmt"
	"os"
	
	log "github.com/sirupsen/logrus"
)

// Server structure.
type server struct {
	ID       string
	Username string
	Password string
}

type repo struct {
	ID        string
	URL       string
	Releases  bool
	Snapshots bool
}

// Settings for the plugin.
type Settings struct {
	UseCentral bool
	Servers    []server
	Repos      []repo
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	// Validation of the settings.
	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	home := "/root"
	user, err := user.Current()
	if err == nil {
		home = user.HomeDir
	}
	settingsPath := path.Join(home, ".m2", "settings.xml")
	settingsFile, err := os.Create(settingsPath)
	if err != nil {
		fmt.Println(err)
		settingsFile.Close()
        return nil;
    }

	log.WithFields(log.Fields{
		"path": settingsPath,
	}).Info("Writing settings.xml")

	tmpl := template.Must(template.ParseFiles("settings.xml"))
	tmpl.Execute(settingsFile, p.settings)
	settingsFile.Close()
	return nil
}
