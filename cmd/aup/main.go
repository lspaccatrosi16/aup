package main

import (
	"fmt"
	"os"

	"github.com/lspaccatrosi16/aup/lib/add"
	"github.com/lspaccatrosi16/aup/lib/remove"
	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/aup/lib/update"
	"github.com/lspaccatrosi16/go-cli-tools/command"
)

func main() {
	cfg := types.Load()

	if len(os.Args) <= 1 {
		interactive(cfg)
	} else {
		flags(cfg)
	}

	types.Save(cfg)
}

func interactive(cfg *types.AUPData) {
	manager := command.NewManager(command.ManagerConfig{Searchable: false})

	manager.Register("add", "Add a new program", provideConfig(cfg, add.Interactive))
	manager.Register("update", "Update a program", provideConfig(cfg, update.Interactive))
	manager.Register("updateall", "Update all programs", provideConfig(cfg, update.InteractiveAll))
	manager.Register("remove", "Remove a program", provideConfig(cfg, remove.Interactive))

	for {
		exit := manager.Tui()
		if exit {
			break
		}
	}

}

func provideConfig(cfg *types.AUPData, f func(*types.AUPData) error) func() error {
	return func() error {
		return f(cfg)
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
