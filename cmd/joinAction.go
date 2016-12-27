package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func joinAction(c *cli.Context) {
	fmt.Println("added task: ", c.Args().First())
}
