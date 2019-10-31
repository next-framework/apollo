package utils

import (
	"os"
	"path/filepath"
)

type FindHelper struct {
	Root     string
	FileName string
}

func (fh *FindHelper) DoFind() error {
	return filepath.Walk(fh.Root, fh.walkHandle)
}

func (fh *FindHelper) walkHandle(path string, info os.FileInfo, err error) error {

}
