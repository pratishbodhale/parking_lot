package main

import (
	"example.com/parking_lot/cli"
	"os"
)

func main() {

	cli := cli.NewCli()
	args := os.Args[1:]

	if len(args) > 0 {
		fileName := args[0]
		cli.ProcessFile(fileName)
	} else {
		cli.ProcessStdIn()
	}
}
