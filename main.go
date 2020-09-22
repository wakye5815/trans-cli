package main

import (
	"fmt"
	"os"

	"example.com.test/command"
)

func main() {
	if err := command.RootCommand.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
