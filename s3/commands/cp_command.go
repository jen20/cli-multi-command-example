package commands

import (
	"github.com/mitchellh/cli"
)

type CPCommand struct {
	Ui cli.Ui
}

func (c *CPCommand) Run(_ []string) int {
	c.Ui.Output("Would run cp here")
	return 0
}

func (c *CPCommand) Help() string {
	return `Copies a local file or S3 object to another location locally or in S3.`
}

func (c *CPCommand) Synopsis() string {
	return "Copies a local file or S3 object to another location locally or in S3."
}

