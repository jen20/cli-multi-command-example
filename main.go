package main

import (
	"fmt"
	"os"

	"github.com/jen20/cli-multi-command-example/ec2"
	"github.com/jen20/cli-multi-command-example/s3"
	"github.com/mitchellh/cli"
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("cli-multi-command-example", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"ec2": func() (cli.Command, error) {
			return &ec2.EC2Command{Ui: ui}, nil
		},
		"s3": func() (cli.Command, error) {
			return &s3.S3Command{Ui: ui}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	os.Exit(exitStatus)
}
