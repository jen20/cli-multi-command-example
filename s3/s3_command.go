package s3

import (
	"github.com/jen20/cli-multi-command-example/s3/commands"
	"github.com/mitchellh/cli"
)

type S3Command struct {
	Ui cli.Ui
}

func (c *S3Command) Run(args []string) int {
	ec2c := cli.NewCLI("cli-multi-command-example s3", "")
	ec2c.Args = args

	ec2c.Commands = map[string]cli.CommandFactory{
		"cp": func() (cli.Command, error) {
			return &commands.CPCommand{Ui: c.Ui}, nil
		},
		"website": func() (cli.Command, error) {
			return &commands.WebsiteCommand{Ui: c.Ui}, nil
		},
	}

	if exitStatus, err := ec2c.Run(); err != nil {
		c.Ui.Error(err.Error())
		return exitStatus
	} else {
		return exitStatus
	}
}

func (c *S3Command) Help() string {
	return "EC2 commands"
}

func (c *S3Command) Synopsis() string {
	return "Commands related to the Simple Storage Service (S3)"
}
