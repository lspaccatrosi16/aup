package repair

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lspaccatrosi16/aup/lib/get"
	"github.com/lspaccatrosi16/aup/lib/types"
)

func Do(cfg *types.AUPData) error {
	toFix := []int{}

	for i, e := range cfg.Entries {
		expPath := cfg.AppPath(e.BinaryName)
		expBinLoc := filepath.Join(expPath, e.BinaryName)

		fs, err := os.Stat(expBinLoc)

		if err != nil {
			if os.IsNotExist(err) {
				toFix = append(toFix, i)
				continue
			} else {
				return err
			}
		}

		if fs.Mode() < 0o755 {
			toFix = append(toFix, i)
			continue
		}
	}

	for _, idx := range toFix {
		ent := cfg.Entries[idx]
		lf, err := get.GHReleaseInfo(ent.RepoKey, ent.ArtifactName)
		if err != nil {
			return err
		}

		err = get.DownloadGHBin(cfg, lf.Url, ent.BinaryName)

		if err != nil {
			return err
		}

		fmt.Printf("Fixed %s\n", ent.ArtDetails())
	}
	return nil

}

func Interactive(cfg *types.AUPData) error {
	return Do(cfg)
}
