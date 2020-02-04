// Copyright (c) 2020, the Drone Maven project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package main

import (
	"encoding/json"

	"github.com/dheerajkhardwal/drone-maven/plugin"
	"github.com/urfave/cli/v2"
)

// Servers structure
type servers struct {
	servers    []plugin.Server
	hasBeenSet bool
}

func (s *servers) String() string {
	return ""
}

func (s *servers) Get() []plugin.Server {
	return s.servers
}

func (s *servers) Set(value string) error {
	if !s.hasBeenSet {
		s.servers = []plugin.Server{}
		s.hasBeenSet = true
	}

	err := json.Unmarshal([]byte(value), &s.servers)
	if err != nil {
		return err
	}
	s.hasBeenSet = true

	return nil
}

// Repos structure
type repos struct {
	repos      []plugin.Repo
	hasBeenSet bool
}

func (r *repos) String() string {
	return ""
}

func (r *repos) Get() []plugin.Repo {
	return r.repos
}

func (r *repos) Set(value string) error {
	if !r.hasBeenSet {
		r.repos = []plugin.Repo{}
		r.hasBeenSet = true
	}

	err := json.Unmarshal([]byte(value), &r.repos)
	if err != nil {
		return err
	}
	r.hasBeenSet = true

	return nil
}

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags() []cli.Flag {
	// Replace below with all the flags required for the plugin.
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "use-central",
			Usage:   "use maven central repository",
			Value:   false,
			EnvVars: []string{"PLUGIN_USE_CENTRAL"},
		},
		&cli.GenericFlag{
			Name:    "servers",
			Usage:   "servers for authentication",
			EnvVars: []string{"PLUGIN_SERVERS"},
			Value:   &servers{},
		},
		&cli.GenericFlag{
			Name:    "repos",
			Usage:   "repsitories to configure",
			EnvVars: []string{"PLUGIN_REPOS"},
			Value:   &repos{},
		},
	}
}

// settingsFromContext creates a plugin.Settings from the cli.Context.
func settingsFromContext(ctx *cli.Context) plugin.Settings {
	return plugin.Settings{
		UseCentral: ctx.Bool("use-central"),
		Servers:    ctx.Generic("servers").(*servers).Get(),
		Repos:      ctx.Generic("repos").(*repos).Get(),
	}
}
