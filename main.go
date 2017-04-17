package main

import (
	"fmt"
	"github.com/urfave/cli"
	"gopkg.in/src-d/go-git.v4"
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
