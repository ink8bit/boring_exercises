// This is a simple CLI app. It has a few commands:
// * ns - retrieve the name servers
// * cname - lookup the CNAME for a given host
// * mx - lookup the mail exchange records for a given host
// * ip - lookup the IP addresses for a given host
package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
