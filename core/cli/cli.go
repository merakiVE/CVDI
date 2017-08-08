package cli

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/merakiVE/CVDI/core/utils"
	"path"
	packageConfig "github.com/merakiVE/CVDI/core/config"
	//"errors"
)

var CommandsCLI cli.Commands
var configGlobal packageConfig.Configuration

func init() {

	configGlobal := packageConfig.Configuration{}
	configGlobal.Load()

	CommandsCLI = cli.Commands{
		{
			Name:        "runserver",
			Aliases:     []string{"run"},
			Category:    "server",
			Usage:       "Run develop server",
			UsageText:   "CVDI run port_number or CVDI runserver",
			Description: "Run develop server",
			Action:      RunServer,
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "Error run server\n")
				return err
			},
		},

		{
			Name:        "generate",
			Aliases:     []string{"gen"},
			Category:    "generator",
			Usage:       "",
			UsageText:   "CVDI generate [subcommand]",
			Description: "Generator",
			ArgsUsage:   "[]",
			Subcommands: cli.Commands{
				cli.Command{
					Name: "keys",
					Description: "Generate public and private keys",
					UsageText: "CVDI generate keys - CVDI gen keys",
					Flags: []cli.Flag{
						cli.BoolTFlag{
							Name:  "force, f",
							Usage: "force generate keys - [warning] replace keys existing",
						},
						cli.BoolTFlag{
							Name:   "exists_keys",
							Hidden: true,
						},
					},
					Before: func(c *cli.Context) error {
						c.Set("exists_keys", "false")

						listErrors := make([]string, 0)
						name_files := []string{"public.key", "public.pem", "private.pem", "private.key"}
						path_keys := configGlobal.GetString("PATH_KEYS")

						for _, name := range name_files {
							path_file := path.Join(path_keys, name)

							if utils.Exists(path_file) {
								listErrors = append(listErrors, fmt.Sprintf("File %s exists", path_file))
							}
						}

						if len(listErrors) > 0 && !c.IsSet("force") {
							c.Set("exists_keys", "true")

							for _, err := range listErrors {
								fmt.Fprintln(c.App.Writer, err)
							}
							fmt.Fprintln(c.App.Writer, "\nUse --force for replace keys existing")
						}

						return nil
					},
					Action: func(c *cli.Context) error {
						path_keys := configGlobal.GetString("PATH_KEYS")
						mesg := "******** Generating public and private keys ********\n"

						if !c.Bool("exists_keys") {
							fmt.Fprintf(c.App.Writer, mesg)
							utils.GenerateKeys(path_keys)
						} else {
							if c.IsSet("force") {
								fmt.Fprintf(c.App.Writer, mesg)
								utils.GenerateKeys(path_keys)
							}
						}
						return nil
					},
					OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
						fmt.Fprintf(c.App.Writer, err.Error())
						return err
					},
				},
			},
			Before: func(c *cli.Context) error {
				path_keys := configGlobal.GetString("PATH_KEYS")

				if utils.IsEmptyString(path_keys) {
					return cli.NewExitError("Not exist key 'PATH_KEYS' in cvdi.conf or the key value is empty", 10)
				}
				return nil
			},
			Action: func(c *cli.Context) {},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, err.Error())
				return err
			},
		},
	}
}
