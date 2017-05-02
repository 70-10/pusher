package main

import (
	"github.com/70-10/pusher/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"push": func() (cli.Command, error) {
			return &command.PushCommand{
				Meta: *meta,
			}, nil
		},
		"config": func() (cli.Command, error) {
			return &command.ConfigCommand{
				Meta: *meta,
			}, nil
		},
		"payload": func() (cli.Command, error) {
			return &command.PayloadCommand{
				Meta: *meta,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
