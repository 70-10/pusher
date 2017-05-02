package command

import (
	"os"
	"os/exec"
)

type PayloadCommand struct {
	Meta
}

func (c *PayloadCommand) Run(args []string) int {
	err := c.Config.InitializePayload()
	if err != nil {
		return ExitCodeInitializePayloadError
	}

	command := "vim" + " " + c.Config.PayloadFilePath

	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return ExitCodeRunCommandError
	}

	return ExitCodeOK
}

func (c *PayloadCommand) Synopsis() string {
	return "Editor payload file"
}

func (c *PayloadCommand) Help() string {
	return ""
}
