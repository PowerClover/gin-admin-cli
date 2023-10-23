package tfs

import (
	"embed"
	"path/filepath"
)

var efsIns embed.FS

func SetEFS(fs embed.FS) {
	efsIns = fs
}

func EFS() embed.FS {
	return efsIns
}

type embedFS struct {
}

func NewEmbedFS() FS {
	return &embedFS{}
}

func (fs *embedFS) ReadFile(name string) ([]byte, error) {
	return efsIns.ReadFile(filepath.Join("tpls", name))
}

func (fs *embedFS) ParseTpl(name string, data interface{}) ([]byte, error) {
	tplBytes, err := fs.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return parseTplData(string(tplBytes), data)
}