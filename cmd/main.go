package main

import (
	"bufio"
	"example.com/parking_lot/cli"
	"fmt"
	"io"
	"os"
	"strings"
)

func ProcessInput(c *cli.Cli, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		command := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if command == "exit"{
			return
		}
		out, err := c.Execute(command)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(out)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func ProcessFile(c *cli.Cli, filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ProcessInput(c, f)
}

func main() {

	cli := cli.NewCli()
	args := os.Args[1:]

	if len(args) > 0 {
		fileName := args[0]
		ProcessFile(cli, fileName)
	} else {
		ProcessInput(cli, os.Stdin)
	}
}
