package update

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lspaccatrosi16/aup/lib/get"
	"github.com/lspaccatrosi16/aup/lib/list"
	"github.com/lspaccatrosi16/aup/lib/types"
)

type UpdateData struct {
	Entry *types.AUPEntry
}

func Gather(cfg *types.AUPData) (*UpdateData, error) {
	opIdx, err := list.GetUserEntryIdx(cfg)
	if err != nil {
		return nil, err
	}
	entChosen := &cfg.Entries[opIdx]

	return &UpdateData{
		Entry: entChosen,
	}, nil
}

func CLI(cfg *types.AUPData) (*UpdateData, error) {
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

	return &UpdateData{
		Entry: eFound,
	}, nil

}

func Do(cfg *types.AUPData, params *UpdateData) error {
	entChosen := params.Entry
	fmt.Println(strings.Repeat("=", 50), "\n", "")
	fmt.Printf("Update %s\n", entChosen.BinDetails())

	latestFile, err := get.GHReleaseInfo(entChosen.RepoKey, entChosen.ArtifactName)

	if err != nil {
		return err
	}

	if latestFile.Version != entChosen.Version {
		err = get.DownloadGHBin(cfg, latestFile.Url, entChosen.BinaryName)
		if err != nil {
			return err
		}
		fmt.Printf("Updated binary %s from %s to %s\n", entChosen.BinaryName, entChosen.Version, latestFile.Version)
		entChosen.Version = latestFile.Version
	} else {
		fmt.Printf("Binary %s is already the latest version\n", entChosen.BinDetails())
	}
	return err
}

func Interactive(cfg *types.AUPData) error {
	params, err := Gather(cfg)
	if err != nil {
		return err
	}
	return Do(cfg, params)
}

func InteractiveAll(cfg *types.AUPData) error {
	for i := 0; i < len(cfg.Entries); i++ {
		e := &cfg.Entries[i]
		params := UpdateData{Entry: e}
		err := Do(cfg, &params)
		if err != nil {
			return err
		}
	}

	return nil
}
