package main

import (
	"fmt"
	"os"

	"github.com/lspaccatrosi16/aup/lib/add"
	"github.com/lspaccatrosi16/aup/lib/configure"
	"github.com/lspaccatrosi16/aup/lib/remove"
	"github.com/lspaccatrosi16/aup/lib/repair"
	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/aup/lib/update"
	"github.com/lspaccatrosi16/aup/lib/version"
	"github.com/lspaccatrosi16/go-cli-tools/command"
)

func main() {
	cfg, err := types.Load()
	if err != nil {
		fmt.Println("error loading config:")
		panic(err)
	}

	if len(os.Args) <= 1 {
		interactive(cfg)
	} else {
		flags(cfg)
	}
}

func interactive(cfg *types.AUPData) {
	manager := command.NewManager(command.ManagerConfig{Searchable: true})

	manager.Register("add", "Add a new program", provideConfig(cfg, add.Interactive))
	manager.Register("update", "Update a program", provideConfig(cfg, update.Interactive))
	manager.Register("updateall", "Update all programs", provideConfig(cfg, update.InteractiveAll))
	manager.Register("remove", "Remove a program", provideConfig(cfg, remove.Interactive))
	manager.Register("version", "See all installed programs", provideConfig(cfg, version.Interactive))
	manager.Register("configure", "Configure AUP", provideConfig(cfg, configure.Interactive))
	manager.Register("repair", "Repair installed programs", provideConfig(cfg, repair.Interactive))

	for {
		exit := manager.Tui()
		types.Save(cfg)
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

	var err error

	switch scmd {
	case "add":
		var params *add.AddData
		params, err = add.CLI()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = add.Do(cfg, params)
	case "update":
		var params *update.UpdateData
		params, err = update.CLI(cfg)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = update.Do(cfg, params)

	case "remove":
		var params *remove.RemoveData
		params, err = remove.CLI(cfg)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = remove.Do(cfg, params)
	case "version":
		version.Do(cfg)
	case "repair":
		err = repair.Do(cfg)
	default:
		fmt.Printf("command \"%s\" is not recognized \n", scmd)
		os.Exit(1)
	}
	if err != nil {
		panic(err)
	}

	types.Save(cfg)
}
