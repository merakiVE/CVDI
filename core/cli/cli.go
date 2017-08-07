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
			UsageText:   "",
			Description: "Generator",
			ArgsUsage:   "[]",
			Subcommands: cli.Commands{
				cli.Command{
					Name: "keys",
					Flags: []cli.Flag{
						cli.BoolTFlag{
							Name:  "force, f",
							Usage: "force generate keys - [warning] replace keys existing",
						},
					},
					Before: func(c *cli.Context) error {
						listErrors := make([]string, 0)
						name_files := []string{"public.key", "public.pem", "private.pem", "private.key"}
						path_keys := configGlobal.GetString("PATH_KEYS")

						if utils.IsEmptyString(path_keys) {
							mesg := "Not exist key 'PATH_KEYS' in cvdi.conf or the key value is empty"
							//return cli.NewExitError(mesg, 10)
							//return errors.New(mesg)
							fmt.Fprintln(c.App.Writer, mesg)
						}

						for _, name := range name_files {
							path_file := path.Join(path_keys, name)

							if utils.Exists(path_file) {
								listErrors = append(listErrors, fmt.Sprintf("File %s exist", path_file))
							}
						}

						if len(listErrors) > 0 && !c.IsSet("force") {
							for _, err := range listErrors {
								fmt.Fprintln(c.App.Writer, err)
							}
							fmt.Fprintln(c.App.Writer, "\nUse --force for replace keys existing")
						}

						return nil
					},
					Action: func(c *cli.Context) error {

						path_keys := configGlobal.GetString("PATH_KEYS")

						if utils.IsEmptyString(path_keys) {
							//return cli.NewExitError("Not exist key 'PATH_KEYS' in cvdi.conf or is empty", 10)
							mesg := "Not exist key 'PATH_KEYS' in cvdi.conf or the key value is empty"
							//return cli.NewExitError(mesg, 10)
							//return errors.New(mesg)
							fmt.Fprintln(c.App.Writer, mesg)
						}

						fmt.Fprintf(c.App.Writer, "******** Generating public and private keys ********\n")

						if c.IsSet("force") {
							fmt.Fprintf(c.App.Writer, "Se ha seteado force\n")
						}

						//utils.GenerateKeys(path_keys)

						return nil
					},
					OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
						fmt.Fprintf(c.App.Writer, err.Error())
						return err
					},
				},
			},
			Action: func(c *cli.Context) {
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, err.Error())
				return err
			},
		},
	}
}
