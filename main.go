package main

import (
	"os"
	"github.com/urfave/cli"
	cliCVDI "github.com/merakiVE/CVDI/core/cli"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "CVDI"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Israel Lugo (MerakiVE)",
			Email: "hostelixisrael@gmail.com, ilugo@bmkeros.org.ve",
		},
	}
	app.Commands = cliCVDI.CommandsCLI
	app.Run(os.Args)
}
