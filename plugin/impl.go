// Copyright (c) 2020, the Drone Maven project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/user"
	"path"
	"text/template"
)

const settingsTemplate = `
<?xml version="1.0" encoding="UTF-8"?>
<settings>
    <servers>
        {{range .Servers}}
        <server>
            <id>{{.ID}}</id>
            <username>{{.Username}}</username>
            <password>{{.Password}}</password>
        </server>
        {{end}}
    </servers>

    <profiles>
        <profile>
            <activation>
                <activeByDefault>true</activeByDefault>
            </activation>
            <repositories>
                {{range .Repos}}
                <repository>
                    <id>{{.ID}}</id>
                    <url>{{.URL}}</url>
                    <layout>default</layout>
                    {{if .Releases}}
                    <releases>
                        <enabled>true</enabled>
                    </releases> 
                    {{end}}
                    {{if .Snapshots}}
                    <snapshots>
                        <enabled>true</enabled>
                    </snapshots>
                    {{end}}
                </repository>
				{{end}}
				{{if .UseCentral}}
                <repository>
                    <id>central</id>
                    <url>https://repo.maven.apache.org/maven2/</url>
                </repository> 
                {{end}}
            </repositories>
        </profile>
    </profiles>
</settings>
`

// Server structure.
type Server struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Repo structure.
type Repo struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Releases  bool   `json:"releases"`
	Snapshots bool   `json:"snapshots"`
}

// Settings for the plugin.
type Settings struct {
	UseCentral bool     `json:"use_central"`
	Servers    []Server `json:"servers"`
	Repos      []Repo   `json:"repos"`
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

	m2Path := path.Join(home, ".m2")
	// Prepare and create .m2 directory if missing.
	os.MkdirAll(m2Path, os.ModePerm)
	settingsPath := path.Join(m2Path, "settings.xml")
	settingsFile, err := os.Create(settingsPath)
	if err != nil {
		fmt.Println(err)
		settingsFile.Close()
		return nil
	}

	log.WithFields(log.Fields{
		"path": settingsPath,
	}).Info("Writing settings.xml")

	tmpl := template.Must(template.New("mvn").Parse(settingsTemplate))
	tmpl.Execute(settingsFile, p.settings)
	settingsFile.Close()
	return nil
}
