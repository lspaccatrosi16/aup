package types

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lspaccatrosi16/go-cli-tools/config"
	"github.com/lspaccatrosi16/go-cli-tools/gbin"
)

type AUPEntry struct {
	BinaryName   string
	ArtifactName string
	Version      string
	RepoKey      string
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

type AUPData struct {
	Entries []AUPEntry
}

func (a *AUPData) String() string {
	buf := bytes.NewBuffer(nil)
	for _, ent := range a.Entries {
		fmt.Fprintln(buf, ent)
	}
	return buf.String()
}

type GHFile struct {
	Name    string
	Version string
	Url     string
}

func Save(a *AUPData) {
	enc := gbin.NewEncoder[AUPData]()
	r, err := enc.EncodeStream(a)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(cfpath())
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, r)
}

func Load() *AUPData {
	f, err := os.Open(cfpath())
	if err != nil {
		if os.IsNotExist(err) {
			return new(AUPData)
		} else {
			panic(err)
		}
	}
	defer f.Close()
	dec := gbin.NewDecoder[AUPData]()
	cfg, err := dec.DecodeStream(f)
	if err != nil {
		panic(err)
	}
	return cfg
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
