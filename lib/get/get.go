package get

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lspaccatrosi16/aup/lib/types"
)

func GHReleaseInfo(repoKey string, artifactName string) (*types.GHFile, error) {
	requestUrl := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repoKey)
	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("code %d: %s", resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()
	respBuff := bytes.NewBuffer(nil)
	io.Copy(respBuff, resp.Body)

	var data ghrel
	err = json.Unmarshal(respBuff.Bytes(), &data)
	if err != nil {
		return nil, err
	}

	var ghFile *types.GHFile

	for _, ast := range data.Assets {
		if ast.Name == artifactName {
			ghFile = &types.GHFile{
				Name:    ast.Name,
				Version: data.TagName,
				Url:     ast.BUrl,
			}
			break
		}
	}

	if ghFile == nil {
		return nil, fmt.Errorf("could not find an artifact that matches name %s", artifactName)
	}

	return ghFile, nil
}

func DownloadGHBin(cfg *types.AUPData, url string, binaryName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	targetPath := filepath.Join(cfg.AppPath(binaryName), binaryName)

	fh, err := os.Create(targetPath)
	if err != nil {
		return err
	}

	io.Copy(fh, resp.Body)

	fh.Close()

	err = os.Chmod(targetPath, 0o755)
	if err != nil {
		return err
	}
	return nil
}
