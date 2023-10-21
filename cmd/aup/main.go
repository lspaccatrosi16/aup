package main

import (
	"fmt"
	"os"

	"github.com/lspaccatrosi16/aup/lib/add"
	"github.com/lspaccatrosi16/aup/lib/remove"
	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/aup/lib/update"
	"github.com/lspaccatrosi16/go-cli-tools/input"
)

func main() {
	cfg := types.Load()

	if len(os.Args) <= 1 {
		interactive(cfg)
	} else {
		flags(cfg)
	}
}

func interactive(cfg *types.AUPData) {
	runoptions := []input.SelectOption{
		{Name: "Add a new program", Value: "a"},
		{Name: "Update a program", Value: "u"},
		{Name: "Remove a program", Value: "r"},
		{Name: "Exit", Value: "e"},
	}

	opt, err := input.GetSelection("App updater", runoptions)
	if err != nil {
		panic(err)
	}

	switch opt {
	case "a":
		params := add.Gather()
		add.Do(cfg, params)
		types.Save(cfg)
		return
	case "u":
		params := update.Gather(cfg)
		update.Do(cfg, params)
		types.Save(cfg)
		return
	case "r":
		params := remove.Gather(cfg)
		remove.Do(cfg, params)
		types.Save(cfg)
		return
	case "e":
		return
	}
}

func flags(cfg *types.AUPData) {
	scmd := os.Args[len(os.Args)-1]

	switch scmd {
	case "add":
		params, err := add.CLI()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		add.Do(cfg, params)
	case "update":
		params, err := update.CLI(cfg)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		update.Do(cfg, params)

	case "remove":
		params, err := remove.CLI(cfg)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		remove.Do(cfg, params)
	default:
		fmt.Printf("command \"%s\" is not recognized \n", scmd)
		os.Exit(1)
	}
}
