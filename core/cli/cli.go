package cli

import (
	"github.com/urfave/cli"
	"fmt"
)

var CommandsCLI cli.Commands

func init() {
	CommandsCLI = cli.Commands{
		{
			Name:        "runserver",
			Aliases:     []string{"run"},
			Category:    "server",
			Usage:       "Run develop server",
			UsageText:   "CVDI run port_number or CVDI runserver",
			Description: "Run develop server",
			//ArgsUsage:   "[arrgh]",
			//Flags: []cli.Flag{
			//	cli.Int64Flag{Name: "port"},
			//},
			//Subcommands: cli.Commands{
			//	cli.Command{
			//		Name:   "wop",
			//		Action: func() {},
			//	},
			//},
			//SkipFlagParsing: false,
			//HideHelp:        false,
			//Hidden:          false,
			//HelpName:        "doo!",
			//BashComplete: func(c *cli.Context) {
			//	fmt.Fprintf(c.App.Writer, "--better\n")
			//},
			//Before: func(c *cli.Context) error {
			//	fmt.Fprintf(c.App.Writer, "brace for impact\n")
			//	return nil
			//},
			//After: func(c *cli.Context) error {
			//	fmt.Fprintf(c.App.Writer, "did we lose anyone?\n")
			//	return nil
			//},
			Action: RunServer,
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "Error run server\n")
				return err
			},
		},
	}
}
