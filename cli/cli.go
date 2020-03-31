package cli

import (
	"bufio"
	"example.com/parking_lot"
	"fmt"
	"io"
	"os"
	"strings"
)

type Cli struct {
	parkingLot parking_lot.ParkingLotManager
}

func NewCli() *Cli{
	c := new(Cli)
	return c
}

func (c *Cli) ProcessInput(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		command := strings.ToLower(scanner.Text())
		if err := c.execute(command); err != nil{
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func (c *Cli) ProcessFile(filePath string){
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	c.ProcessInput(f)
}

func (c *Cli) ProcessStdIn(){
	c.ProcessInput(os.Stdin)
}