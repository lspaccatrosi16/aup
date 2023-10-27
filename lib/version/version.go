package version

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/lspaccatrosi16/aup/lib/types"
)

const div = "  |  "

func Do(cfg *types.AUPData) {
	if len(cfg.Entries) == 0 {
		fmt.Println("There are no programs currently installed")
		return
	}

	buf := bytes.NewBuffer(nil)
	names := []string{}
	versions := []string{}

	maxName := 0
	maxVersion := 0

	for _, e := range cfg.Entries {
		str := e.ArtDetails()
		if len(str) > maxName {
			maxName = len(str)
		}

		if len(e.Version) > maxVersion {
			maxVersion = len(e.Version)
		}

		names = append(names, str)
		versions = append(versions, e.Version)
	}

	bar := strings.Repeat("=", maxName+maxVersion+len(div))

	fmt.Fprintln(buf, "Currently installed programs")
	fmt.Fprintln(buf, bar)

	for i := 0; i < len(cfg.Entries); i++ {
		name := names[i]
		version := versions[i]
		fmt.Fprintf(buf, "%-*s%s%s\n", maxName, name, div, version)
	}

	fmt.Fprintln(buf, bar)

	fmt.Println(buf.String())
}

func Interactive(cfg *types.AUPData) error {
	Do(cfg)
	return nil
}
