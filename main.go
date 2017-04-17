package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hoge"
	app.Action = func(c *cli.Context) error {
		fmt.Println("hoge")
		return nil
	}

	app.Run(os.Args)
}
