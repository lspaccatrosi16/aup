package remove

import (
	"github.com/lspaccatrosi16/aup/lib/list"
	"github.com/lspaccatrosi16/aup/lib/types"
)

func Remove(cfg *types.AUPData) {
	delIdx := list.GetUserEntryIdx(cfg)

	newEnts := []types.AUPEntry{}

	for i, ent := range cfg.Entries {
		if i == delIdx {
			continue
		} else {
			newEnts = append(newEnts, ent)
		}
	}

	cfg.Entries = newEnts
}
