package remove

import (
	"flag"
	"fmt"

	"github.com/lspaccatrosi16/aup/lib/list"
	"github.com/lspaccatrosi16/aup/lib/types"
)

type removeData struct {
	EntryIdx int
}

func Gather(cfg *types.AUPData) *removeData {
	delIdx := list.GetUserEntryIdx(cfg)
	return &removeData{
		EntryIdx: delIdx,
	}
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
