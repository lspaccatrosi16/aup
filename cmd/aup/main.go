package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/lspaccatrosi16/aup/lib/add"
	"github.com/lspaccatrosi16/aup/lib/configure"
	"github.com/lspaccatrosi16/aup/lib/remove"
	"github.com/lspaccatrosi16/aup/lib/repair"
	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/aup/lib/update"
	"github.com/lspaccatrosi16/aup/lib/version"
	"github.com/lspaccatrosi16/go-cli-tools/command"
)

var help = flag.Bool("h", false, "shows help message")
var repokey = flag.String("r", "", "repoKey of the executable")
var artifactName = flag.String("a", "", "name of the artifact")
var binaryName = flag.String("b", "", "local name of the binary to use")

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

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

type commandReq struct {
	CmdName string
	Args    []bool
}

var commands = []commandReq{
	{"add", []bool{true, true, true}},
	{"update", []bool{true, false, false}},
	{"remove", []bool{true, false, false}},
	{"version", []bool{false, false, false}},
	{"repair", []bool{false, false, false}},
	{"help", []bool{false, false, false}},
}

var reqFlagName = []string{"a", "b", "r"}

func flags(cfg *types.AUPData) {
	scmd := flag.Args()[0]

	found := false
	params := []string{*artifactName, *binaryName, *repokey}

	for _, c := range commands {
		if scmd == c.CmdName {
			found = true
			for i, b := range c.Args {
				if b && params[i] == "" {
					fmt.Printf("invalid arguments: expected -%s\n", reqFlagName[i])
					os.Exit(1)
				}
			}
		}
	}

	if !found {
		fmt.Printf("command \"%s\" was not recognized \n", scmd)
		scmd = "help"
	}

	var err error

	switch scmd {
	case "add":
		var params *add.AddData
		params, err = add.CLI(*artifactName, *binaryName, *repokey)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = add.Do(cfg, params)
	case "update":
		var params *update.UpdateData
		params, err = update.CLI(cfg, *artifactName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = update.Do(cfg, params)

	case "remove":
		var params *remove.RemoveData
		params, err = remove.CLI(cfg, *artifactName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = remove.Do(cfg, params)
	case "version":
		version.Do(cfg)
	case "repair":
		err = repair.Do(cfg)
	case "help":
		fmt.Println("Availible Commands:")

		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		for _, c := range commands {
			reqStr := ""
			for i, b := range c.Args {
				if b {
					reqStr += " -" + reqFlagName[i]
				}
			}
			fmt.Fprintf(writer, "%s\t%s\n", c.CmdName, reqStr)
		}

		writer.Flush()
		os.Exit(0)
	}
	if err != nil {
		panic(err)
	}

	types.Save(cfg)
}
