package add

import (
	"fmt"
	"strings"

	"github.com/lspaccatrosi16/aup/lib/get"
	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/go-cli-tools/input"
)

func Add(cfg *types.AUPData) {
	validator := func(str string) error {
		splitted := strings.Split(str, "/")
		if len(splitted) != 2 {
			return fmt.Errorf("repokey must be formatted in user/repo format e.g. lspaccatrosi16/scaffold")
		}

		if len(splitted[0]) == 0 {
			return fmt.Errorf("user must not be an empty string")
		}

		if len(splitted[1]) == 0 {
			return fmt.Errorf("repo must not be an empty string")
		}

		return nil
	}
	repoKey := input.GetValidatedInput("Repokey", validator)
	artifactName := input.GetInput("Artifact Name")
	binaryName := input.GetInput("Binary name (leave blank for artifact name)")
	if binaryName == "" {
		binaryName = artifactName
	}

	file, err := get.GetGHFile(repoKey, artifactName)
	if err != nil {
		panic(err)
	}
	entry := types.AUPEntry{
		BinaryName:   binaryName,
		ArtifactName: artifactName,
		RepoKey:      repoKey,
		Version:      file.Version,
	}

	cfg.Entries = append(cfg.Entries, entry)
	get.DGHFile(file.Url, binaryName)

	fmt.Printf("Got binary %s@%s\n", entry.BinaryName, entry.Version)
}
