package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "gbr"
	app.Action = func(c *cli.Context) error {
		if !checkGit() {
			println("not git dir")
		}

		checkLocalBranches()
		// 		checkRemoteBranches()
		return nil
	}

	app.Run(os.Args)
}

func checkGit() bool {

	dir, _ := os.Getwd()
	println(string(dir))
	return false
}

func checkLocalBranches() {
	println("check local branches")
	out, _ := exec.Command("git", "branch").Output()
	result := string(out)
	slice := strings.Split(result, "\r")
	for _, x := range slice {
		println(x)
	}
}

func checkRemoteBranches() {
	println("check remote branches")
	out, _ := exec.Command("git", "branch", "-a").Output()
	result := string(out)
	slice := strings.Split(result, "\r")
	for _, x := range slice {
		println(x)
	}
}
