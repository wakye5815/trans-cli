package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/wakye5815/trans-cli/command"
)

func main() {
	viper.SetConfigFile("./.env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("file not found")
		} else {
			panic(err)
		}
	}

	if err := command.NewRootCommand().Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
