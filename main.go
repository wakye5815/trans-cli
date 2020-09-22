package main

import (
	"fmt"
	"os"

	"github.com/wakye5815/trans-cli/command"
)

func main() {
	if err := command.RootCommand.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
