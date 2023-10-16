package list

import (
	"fmt"

	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/go-cli-tools/input"
)

func GetUserEntryIdx(cfg *types.AUPData) int {
	options := []input.SelectOption{}

	for _, ent := range cfg.Entries {
		options = append(options, input.SelectOption{
			Name:  fmt.Sprintf("%s: %s", ent.RepoKey, ent.ArtifactName),
			Value: "",
		})
	}

	_, opIdx, err := input.GetSelectionIdx("Choose a Binary:", options)
	if err != nil {
		panic(err)
	}

	return opIdx

}
