package types

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lspaccatrosi16/go-cli-tools/config"
	"github.com/lspaccatrosi16/go-libs/gbin"
)

type AUPEntry struct {
	BinaryName   string
	ArtifactName string
	Version      string
	RepoKey      string
}

func (a *AUPEntry) ArtDetails() string {
	return fmt.Sprintf("%s/%s", a.RepoKey, a.ArtifactName)
}

func (a *AUPEntry) BinDetails() string {
	return fmt.Sprintf("%s@%s", a.BinaryName, a.Version)
}

func (a *AUPEntry) String() string {
	buf := bytes.NewBuffer(nil)

	fmt.Fprintln(buf, strings.Repeat("=", 50))
	fmt.Fprintf(buf, "Binary Name    : %s\n", a.BinaryName)
	fmt.Fprintf(buf, "Artifact Name  : %s\n", a.ArtifactName)
	fmt.Fprintf(buf, "Latest Version : %s\n", a.Version)
	fmt.Fprintf(buf, "RepoKey        : %s\n", a.RepoKey)

	return buf.String()
}

type AUPConfig struct {
	AppSavePath string
}

type AUPData struct {
	Entries []AUPEntry
	Config  *AUPConfig
}

func (a *AUPData) String() string {
	buf := bytes.NewBuffer(nil)
	for _, ent := range a.Entries {
		fmt.Fprintln(buf, ent)
	}
	return buf.String()
}

func (a *AUPData) AppPath(bin string) string {
	path := ""
	if a.Config.AppSavePath != "" {
		path = filepath.Join(a.Config.AppSavePath, bin)
	} else {
		path = filepath.Join(CPath(), "apps", bin)
	}

	err := os.MkdirAll(path, 0o755)
	if err != nil {
		panic(err)
	}

	return path
}

type GHFile struct {
	Name    string
	Version string
	Url     string
}

func Save(a *AUPData) error {
	buf := bytes.NewBuffer(nil)
	for _, e := range a.Entries {
		dir := a.AppPath(e.BinaryName)
		fmt.Fprintf(buf, "export PATH=\"$PATH:%s\"\n", dir)
	}

	pathFile := filepath.Join(CPath(), ".auprc")
	f, err := os.Create(pathFile)
	if err != nil {
		return err
	}

	io.Copy(f, buf)
	f.Close()

	enc := gbin.NewEncoder[AUPData]()
	r, err := enc.EncodeStream(a)
	if err != nil {
		return err
	}

	f, err = os.Create(cfpath())
	if err != nil {
		return err
	}

	io.Copy(f, r)
	f.Close()
	return nil
}

func Load() (*AUPData, error) {
	f, err := os.Open(cfpath())
	if err != nil {
		if os.IsNotExist(err) {
			d := new(AUPData)
			d.Config = new(AUPConfig)
			return d, nil
		} else {
			return nil, err
		}
	}
	defer f.Close()
	dec := gbin.NewDecoder[AUPData]()
	cfg, err := dec.DecodeStream(f)

	if cfg.Config == nil {
		cfg.Config = new(AUPConfig)
	}

	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func cfpath() string {
	return filepath.Join(CPath(), "cfg.bin")
}

func CPath() string {
	p, err := config.GetConfigPath("aup")
	if err != nil {
		panic(err)
	}
	return p
}
