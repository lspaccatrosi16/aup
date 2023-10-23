package update

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lspaccatrosi16/aup/lib/get"
	"github.com/lspaccatrosi16/aup/lib/list"
	"github.com/lspaccatrosi16/aup/lib/types"
)

type updateData struct {
	Entry *types.AUPEntry
}

func Gather(cfg *types.AUPData) (*updateData, error) {
	opIdx, err := list.GetUserEntryIdx(cfg)
	if err != nil {
		return nil, err
	}
	entChosen := &cfg.Entries[opIdx]

	return &updateData{
		Entry: entChosen,
	}, nil
}

func CLI(cfg *types.AUPData) (*updateData, error) {
	var artifactName string
	flag.StringVar(&artifactName, "a", "", "name of the artifact")

	flag.Parse()

	if artifactName == "" {
		return nil, fmt.Errorf("artifact name must not be \"\"")
	}

	var eFound *types.AUPEntry

	for i := 0; i < len(cfg.Entries); i++ {
		ent := cfg.Entries[i]
		if ent.ArtifactName == artifactName {
			eFound = &(ent)
		}
	}

	if eFound == nil {
		return nil, fmt.Errorf("could not find a binary with artifact name \"%s\"", artifactName)
	}

	return &updateData{
		Entry: eFound,
	}, nil

}

func Do(cfg *types.AUPData, params *updateData) {
	entChosen := params.Entry
	fmt.Println(strings.Repeat("=", 50), "\n", "")
	fmt.Printf("Update %s@%s\n", entChosen.ArtifactName, entChosen.Version)

	latestFile, err := get.GetGHFile(entChosen.RepoKey, entChosen.ArtifactName)

	if err != nil {
		panic(err)
	}

	if latestFile.Version != entChosen.Version {
		get.DGHFile(latestFile.Url, entChosen.BinaryName)
		fmt.Printf("Updated binary %s from %s to %s\n", entChosen.BinaryName, entChosen.Version, latestFile.Version)
		entChosen.Version = latestFile.Version
	} else {
		fmt.Printf("Binary %s@%s is already the latest version\n", entChosen.BinaryName, entChosen.Version)
	}

}

func Interactive(cfg *types.AUPData) error {
	params, err := Gather(cfg)
	if err != nil {
		return err
	}
	Do(cfg, params)
	return nil
}

func InteractiveAll(cfg *types.AUPData) error {
	for _, e := range cfg.Entries {
		params := updateData{Entry: &e}
		Do(cfg, &params)
	}
	return nil
}
