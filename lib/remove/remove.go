package remove

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lspaccatrosi16/aup/lib/list"
	"github.com/lspaccatrosi16/aup/lib/types"
)

type removeData struct {
	EntryIdx int
}

func Gather(cfg *types.AUPData) (*removeData, error) {
	delIdx, err := list.GetUserEntryIdx(cfg)
	if err != nil {
		return nil, err
	}

	return &removeData{
		EntryIdx: delIdx,
	}, nil
}

func CLI(cfg *types.AUPData) (*removeData, error) {
	var artifactName string
	flag.StringVar(&artifactName, "a", "", "name of the artifact")

	flag.Parse()

	if artifactName == "" {
		return nil, fmt.Errorf("artifact name must not be \"\"")
	}

	iFound := -1

	for i := 0; i < len(cfg.Entries); i++ {
		ent := cfg.Entries[i]
		if ent.ArtifactName == artifactName {
			iFound = i
		}
	}

	if iFound == -1 {
		return nil, fmt.Errorf("could not find a binary with artifact name \"%s\"", artifactName)
	}

	return &removeData{
		EntryIdx: iFound,
	}, nil

}

func Do(cfg *types.AUPData, params *removeData) {
	entRemove := cfg.Entries[params.EntryIdx]
	fmt.Printf("\nRemove %s@%s\n", entRemove.BinaryName, entRemove.Version)
	fmt.Println(strings.Repeat("=", 50))

	newEnts := []types.AUPEntry{}
	for i, ent := range cfg.Entries {
		if i == params.EntryIdx {
			continue
		} else {
			newEnts = append(newEnts, ent)
		}
	}

	cfg.Entries = newEnts
}

func Interactive(cfg *types.AUPData) error {
	params, err := Gather(cfg)
	if err != nil {
		return err
	}

	Do(cfg, params)
	return nil
}
