package list

import (
	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/go-cli-tools/command"
)

func GetUserEntryIdx(cfg *types.AUPData) (int, error) {
	manager := command.NewManager(command.ManagerConfig{Searchable: true})

	for i, ent := range cfg.Entries {
		manager.RegisterData(ent.ArtifactName, ent.RepoKey, encloseIndex(i))
	}

	vAny, err := manager.DataTui()

	if err != nil {
		return 0, err
	}

	idx := vAny.(int)

	return idx, nil
}

func encloseIndex(idx int) func() (any, error) {
	return func() (any, error) {
		return idx, nil
	}
}
