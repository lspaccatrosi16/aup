package update

import (
	"fmt"

	"github.com/lspaccatrosi16/aup/lib/get"
	"github.com/lspaccatrosi16/aup/lib/list"
	"github.com/lspaccatrosi16/aup/lib/types"
)

func Update(cfg *types.AUPData) {
	opIdx := list.GetUserEntryIdx(cfg)
	entChosen := &cfg.Entries[opIdx]

	latestFile, err := get.GetGHFile(entChosen.RepoKey, entChosen.ArtifactName)

	if err != nil {
		panic(err)
	}

	if latestFile.Version != entChosen.Version {
		get.DGHFile(latestFile.Url, entChosen.BinaryName)
		fmt.Printf("Updated binary %s from %s to %s\n", entChosen.BinaryName, entChosen.Version, latestFile.Version)
		entChosen.Version = latestFile.Version
	} else {
		fmt.Printf("Binary %s@%s is the latest version\n", entChosen.BinaryName, entChosen.Version)
	}
}
