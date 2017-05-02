package command

import (
	"os"
	"os/exec"
	"path/filepath"
)

type ConfigCommand struct {
	Meta
}

func (c *ConfigCommand) Run(args []string) int {
	command := "vim" + " " + filepath.Join(os.Getenv("HOME"), ".config", "pusher", "config.toml")
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		return ExitCodeRunCommandError
	}

	return ExitCodeOK
}

func (c *ConfigCommand) Synopsis() string {
	return "Editor config file"
}

func (c *ConfigCommand) Help() string {
	return ""
}
