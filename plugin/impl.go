// Copyright (c) 2020, the Drone Maven project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"errors"
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
				{{if .Central}}
                <repository>
                    <id>central</id>
                    <url>{{.CentralRepo}}</url>
                </repository> 
                {{end}}
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
						<updatePolicy>always</updatePolicy>
                    </snapshots>
                    {{end}}
                </repository>
				{{end}}
            </repositories>
        </profile>
    </profiles>
</settings>
`

// Server structure.
type server struct {
	ID       string
	Username string
	Password string
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
	Username    string `json:"username"`
	Password    string `json:"password"`
	Repos       []Repo `json:"repos"`
	Central     bool   `json:"central"`
	CentralRepo string `json:"central_repo"`
}

type internalSettings struct {
	Servers     []server
	Repos       []Repo
	Central     bool
	CentralRepo string
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	// Validation of the settings.
	settings := p.settings
	log.WithFields(log.Fields{
		"enabled": settings.Central,
		"repo":    settings.CentralRepo,
	}).Debug("Configuration for central repository")

	repos := settings.Repos
	if len(repos) == 0 {
		log.Warn("No repo configured in settings")
	} else {
		for index := range repos {
			repo := repos[index]
			if repo.ID == "" {
				log.Error("`id` field missing for repository entry")
				return errors.New("missing 'id' field in repository entry")
			}
			if repo.URL == "" {
				log.Error("`url` field missing for repository entry")
				return errors.New("missing 'url' field in repository entry")
			}
		}
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	home := "/root"
	user, err := user.Current()
	if err == nil {
		home = user.HomeDir
	}
	fmt.Printf("Resolved home directory: %s", home)

	m2Path := path.Join(home, ".m2")
	// Prepare and create .m2 directory if missing.
	os.MkdirAll(m2Path, os.ModePerm)
	settingsPath := path.Join(m2Path, "settings.xml")
	fmt.Printf("Settings file path: %s\n", settingsPath)
	settingsFile, err := os.Create(settingsPath)
	if err != nil {
		settingsFile.Close()
		fmt.Println(err)
		return err
	}

	repos := p.settings.Repos
	username := p.settings.Username
	password := p.settings.Password
	// Transform p.settings to internalSettings.
	settings := &internalSettings{}
	settings.Repos = repos
	settings.Central = p.settings.Central
	settings.CentralRepo = p.settings.CentralRepo
	settings.Servers = make([]server, len(repos)) // Initialize servers
	// For each repo, create a server with given repo ID
	for i := 0; i < len(repos); i++ {
		settings.Servers[i] = server{repos[i].ID, username, password}
	}

	tmpl := template.Must(template.New("mvn").Parse(settingsTemplate))
	tmpl.Execute(settingsFile, settings)
	settingsFile.Close()
	fmt.Println("Created settings file successfully.")
	return nil
}
