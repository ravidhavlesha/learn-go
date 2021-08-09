package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/ravidhavlesha/learn-go/webcli/actions"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func info() {
	app.Name = "Website Lookup CLI"
	app.Usage = "Query IPs, CNAMEs, MX Records and Name Servers!"
	app.Version = "1.0.0"
}

func setup() {
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(c.App.Writer, "Error: %v\n", err)
		return nil
	}
}

func flags() {
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "host",
			Aliases: []string{"l"},
		},
	}
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:         "nameserver",
			Aliases:      []string{"ns"},
			Flags:        app.Flags,
			Usage:        "Name server for the particular host",
			OnUsageError: app.OnUsageError,
			Action:       actions.NameServers,
		},
		{
			Name:         "ipaddress",
			Aliases:      []string{"ip"},
			Flags:        app.Flags,
			Usage:        "IP address for the particular host",
			OnUsageError: app.OnUsageError,
			Action:       actions.IPAddress,
		},
		{
			Name:         "cname",
			Aliases:      []string{"cn"},
			Flags:        app.Flags,
			Usage:        "CNAME for the particular host",
			OnUsageError: app.OnUsageError,
			Action:       actions.CNAME,
		},
		{
			Name:         "mxrecords",
			Aliases:      []string{"mx"},
			Flags:        app.Flags,
			Usage:        "MX records for the particular host",
			OnUsageError: app.OnUsageError,
			Action:       actions.MXRecords,
		},
		{
			Name:         "all",
			Aliases:      []string{"a"},
			Flags:        app.Flags,
			Usage:        "All data for the particular host",
			OnUsageError: app.OnUsageError,
			Action: func(c *cli.Context) error {
				fmt.Println("Name Servers:")
				actions.NameServers(c)
				fmt.Println("")

				fmt.Println("IP Address:")
				actions.IPAddress(c)
				fmt.Println("")

				fmt.Println("CNAME:")
				actions.CNAME(c)
				fmt.Println("")

				fmt.Println("MX Records:")
				actions.MXRecords(c)

				return nil
			},
		},
	}

}

func main() {
	info()
	setup()
	flags()
	commands()
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
