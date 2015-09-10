package commands

import (
	"github.com/mitchellh/cli"
)

type DescribeInstancesCommand struct {
	Ui cli.Ui
}

func (c *DescribeInstancesCommand) Run(_ []string) int {
	c.Ui.Output("Would run describe-instances here")
	return 0
}

func (c *DescribeInstancesCommand) Help() string {
	return `Describes one or more of your instances.

If you specify one or more instance IDs, Amazon EC2 returns information for those instances. If you do not specify instance IDs, Amazon EC2 returns information for all relevant instances. If you specify an instance ID that is not valid, an error is returned. If you specify an instance that you do not own, it is not included in the returned results.

Recently terminated instances might appear in the returned results. This interval is usually less than one hour.`
}

func (c *DescribeInstancesCommand) Synopsis() string {
	return "Describes one or more of your instances."
}
