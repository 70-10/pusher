package command

import (
	"github.com/70-10/pusher/config"
	"github.com/mitchellh/cli"
)

type Meta struct {
	Ui     cli.Ui
	Config config.Config
}
