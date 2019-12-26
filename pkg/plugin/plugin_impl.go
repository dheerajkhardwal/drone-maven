// Copyright (c) 2019, the Drone Plugins project authors.
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

// Settings for the Plugin.
type Settings struct {
	// Fill in the data structure with appropriate values
	UseCentral bool
	Servers    []server
	Repos      []repo
}

func (p *pluginImpl) Validate() error {
	// Validate the Config and return an error if there are issues.
	return nil
}

func (p *pluginImpl) Exec() error {
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
