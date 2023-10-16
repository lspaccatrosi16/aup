package main

import (
	"github.com/lspaccatrosi16/aup/lib/add"
	"github.com/lspaccatrosi16/aup/lib/remove"
	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/aup/lib/update"
	"github.com/lspaccatrosi16/go-cli-tools/input"
)

func main() {
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

	cfg := types.Load()

	switch opt {
	case "a":
		add.Add(cfg)
		types.Save(cfg)
		return
	case "u":
		update.Update(cfg)
		types.Save(cfg)
		return
	case "r":
		remove.Remove(cfg)
		types.Save(cfg)
		return
	case "e":
		return
	}
}
