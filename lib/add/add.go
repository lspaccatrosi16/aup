package add

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lspaccatrosi16/aup/lib/get"
	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/go-cli-tools/input"
)

type addData struct {
	RKey  string
	AName string
	BName string
}

func validator(str string) error {
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

func Gather() *addData {
	fmt.Println("New Program")
	fmt.Println(strings.Repeat("=", 50))
	repoKey := input.GetValidatedInput("Repokey", validator)
	artifactName := input.GetInput("Artifact Name")
	binaryName := input.GetInput("Binary name (leave blank for artifact name)")
	if binaryName == "" {
		binaryName = artifactName
	}

	return &addData{
		RKey:  repoKey,
		AName: artifactName,
		BName: binaryName,
	}
}

func CLI() (*addData, error) {
	var repoKey string
	var artifactName string
	var binaryName string

	flag.StringVar(&repoKey, "r", "", "repoKey of the executable")
	flag.StringVar(&artifactName, "a", "", "name of the artifact")
	flag.StringVar(&binaryName, "b", "", "local name of the binary to use")

	flag.Parse()

	if artifactName == "" {
		return nil, fmt.Errorf("artifact name must not be \"\"")
	}

	if binaryName == "" {
		binaryName = artifactName
	}

	verr := validator(repoKey)

	if verr != nil {
		return nil, verr
	}

	return &addData{
		RKey:  repoKey,
		AName: artifactName,
		BName: binaryName,
	}, nil

}

func Do(cfg *types.AUPData, params *addData) {
	fmt.Println(strings.Repeat("=", 50), "\n", "")
	fmt.Printf("Searching for %s/%s\n", params.RKey, params.AName)

	file, err := get.GetGHFile(params.RKey, params.AName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Get %s@%s\n", params.AName, file.Version)

	entry := types.AUPEntry{
		BinaryName:   params.BName,
		ArtifactName: params.AName,
		RepoKey:      params.RKey,
		Version:      file.Version,
	}

	cfg.Entries = append(cfg.Entries, entry)
	get.DGHFile(file.Url, params.BName)

	fmt.Printf("Got binary %s@%s\n", entry.BinaryName, entry.Version)

}

func Interactive(cfg *types.AUPData) error {
	params := Gather()
	Do(cfg, params)
	return nil
}
