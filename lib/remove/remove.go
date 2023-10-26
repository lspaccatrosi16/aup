package remove

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

func Do(cfg *types.AUPData, params *removeData) error {
	entRemove := cfg.Entries[params.EntryIdx]
	fmt.Println(strings.Repeat("=", 50), "\n", "")
	fmt.Printf("Remove %s@%s\n", entRemove.ArtifactName, entRemove.Version)

	newEnts := []types.AUPEntry{}
	for i, ent := range cfg.Entries {
		if i == params.EntryIdx {
			continue
		} else {
			newEnts = append(newEnts, ent)
		}
	}

	path := filepath.Join(cfg.AppPath(entRemove.BinaryName), entRemove.BinaryName)

	err := os.Remove(path)

	if err != nil {
		return err
	}

	cfg.Entries = newEnts
	return nil
}

func Interactive(cfg *types.AUPData) error {
	params, err := Gather(cfg)
	if err != nil {
		return err
	}

	err = Do(cfg, params)

	if err != nil {
		return err
	}
	return nil
}
