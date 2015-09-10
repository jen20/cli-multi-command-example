package commands

import (
	"github.com/mitchellh/cli"
	"fmt"
)

type CreatePlacementGroupCommand struct {
	Ui cli.Ui
}

func (c *CreatePlacementGroupCommand) Run(args []string) int {
	c.Ui.Output("Would run create-placement-group here")
	c.Ui.Output(fmt.Sprintf("%+v", args))
	return 0
}

func (c *CreatePlacementGroupCommand) Help() string {
	return `Describes one or more of your instances.

If you specify one or more instance IDs, Amazon EC2 returns information for those instances. If you do not specify instance IDs, Amazon EC2 returns information for all relevant instances. If you specify an instance ID that is not valid, an error is returned. If you specify an instance that you do not own, it is not included in the returned results.

Recently terminated instances might appear in the returned results. This interval is usually less than one hour.`
}

func (c *CreatePlacementGroupCommand) Synopsis() string {
	return "Creates a placement group that you launch cluster instances into."
}
