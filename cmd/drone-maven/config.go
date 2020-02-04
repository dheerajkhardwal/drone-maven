// Copyright (c) 2020, the Drone Maven project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/dheerajkhardwal/drone-maven/plugin"
	"github.com/urfave/cli/v2"
)

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
		fmt.Println(err)
		return err
	}
	r.hasBeenSet = true

	return nil
}

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags() []cli.Flag {
	// Replace below with all the flags required for the plugin.
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "username",
			Usage:   "username for maven repository",
			EnvVars: []string{"PLUGIN_USERNAME"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    "password",
			Usage:   "password for maven repository",
			EnvVars: []string{"PLUGIN_PASSWORD"},
			Required: true,
		},
		&cli.GenericFlag{
			Name:    "repos",
			Usage:   "repsitories to configure",
			EnvVars: []string{"PLUGIN_REPOS"},
			Value:   &repos{},
		},
		&cli.BoolFlag{
			Name:    "central",
			Usage:   "enable maven central repository",
			Value:   false,
			EnvVars: []string{"PLUGIN_CENTRAL"},
		},
		&cli.StringFlag{
			Name:    "central_repo",
			Usage:   "maven central repository",
			Value:   "https://repo.maven.apache.org/maven2/",
			EnvVars: []string{"PLUGIN_CENTRAL_REPO"},
		},
	}
}

// settingsFromContext creates a plugin.Settings from the cli.Context.
func settingsFromContext(ctx *cli.Context) plugin.Settings {
	return plugin.Settings{
		Username: ctx.String("username"),
		Password: ctx.String("password"),
		Repos:    ctx.Generic("repos").(*repos).Get(),
		Central:  ctx.Bool("central"),
		CentralRepo: ctx.String("central_repo"),
	}
}
