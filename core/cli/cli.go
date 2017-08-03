package cli

import (
	"github.com/urfave/cli"
	"fmt"
)

var CommandsCLI cli.Commands

func init() {
	CommandsCLI = cli.Commands{
		{
			Name:    "runserver",
			Aliases: []string{"run"},
			Usage:   "run develop server",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		/*{
			Name:        "runserver",
			Aliases:     []string{"run"},
			Category:    "server",
			Usage:       "run develop server",
			UsageText:   "run develop server",
			Description: "run develop server",
			ArgsUsage:   "[arrgh]",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "forever, forevvarr"},
			},
			Subcommands: cli.Commands{
				cli.Command{
					Name:   "wop",
					Action: func() {},
				},
			},
			SkipFlagParsing: false,
			HideHelp:        false,
			Hidden:          false,
			HelpName:        "doo!",
			BashComplete: func(c *cli.Context) {
				fmt.Fprintf(c.App.Writer, "--better\n")
			},
			Before: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "brace for impact\n")
				return nil
			},
			After: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "did we lose anyone?\n")
				return nil
			},
			Action: func(c *cli.Context) error {
				c.Command.FullName()
				c.Command.HasName("wop")
				c.Command.Names()
				c.Command.VisibleFlags()
				fmt.Fprintf(c.App.Writer, "dodododododoodododddooooododododooo\n")
				if c.Bool("forever") {
					c.Command.Run(c)
				}
				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "for shame\n")
				return err
			},
		},*/
	}
}
