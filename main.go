// This is a simple CLI app. It has a few commands:
// * ns - retrieve the name servers
// * cname - lookup the CNAME for a given host
// * mx - lookup the mail exchange records for a given host
// * ip - lookup the IP addresses for a given host
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	version = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "An example CLI app"
	app.Usage = "Shows you a greeting message"
	app.Version = version

	// app.Flags = []cli.Flag{
	// 	cli.StringFlag{
	// 		Name:  "hello, e",
	// 		Value: "Hello stranger",
	// 		Usage: "Shows hello message",
	// 	},
	// }

	app.Commands = []cli.Command{
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "print an array",
			Action: func(c *cli.Context) error {
				arr := [5]int{1, 2, 3, 4, 5}
				fmt.Println(arr)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
