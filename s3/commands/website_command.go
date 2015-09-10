package commands

import (
	"github.com/mitchellh/cli"
)

type WebsiteCommand struct {
	Ui cli.Ui
}

func (c *WebsiteCommand) Run(_ []string) int {
	c.Ui.Output("Would run website here")
	return 0
}

func (c *WebsiteCommand) Help() string {
	return `Set the website configuration for a bucket.`
}

func (c *WebsiteCommand) Synopsis() string {
	return "Set the website configuration for a bucket."
}
