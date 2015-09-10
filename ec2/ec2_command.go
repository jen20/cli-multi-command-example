package ec2

import (
	"github.com/jen20/cli-multi-command-example/ec2/commands"
	"github.com/mitchellh/cli"
)

type EC2Command struct {
	Ui cli.Ui
}

func (c *EC2Command) Run(args []string) int {
	ec2c := cli.NewCLI("cli-multi-command-example ec2", "")
	ec2c.Args = args

	ec2c.Commands = map[string]cli.CommandFactory{
		"create-placement-group": func() (cli.Command, error) {
			return &commands.CreatePlacementGroupCommand{Ui: c.Ui}, nil
		},
		"describe-instances": func() (cli.Command, error) {
			return &commands.DescribeInstancesCommand{Ui: c.Ui}, nil
		},
	}

	if exitStatus, err := ec2c.Run(); err != nil {
		c.Ui.Error(err.Error())
		return exitStatus
	} else {
		return exitStatus
	}
}

func (c *EC2Command) Help() string {
	return "EC2 commands"
}

func (c *EC2Command) Synopsis() string {
	return "Commands related to the Elastic Compute Cloud (EC2)"
}
