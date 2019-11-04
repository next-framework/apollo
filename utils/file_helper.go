package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func GetSuffix(filename string) string {
	onlyFilenameWithSuffix := path.Base(filename)
	suffix := path.Ext(onlyFilenameWithSuffix)
	return suffix
}

type JudgeFunc func(pattern, filename string) bool

func Find(path, pattern string, judgeFunc JudgeFunc) (string, error) {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}

	sep := string(os.PathSeparator)

	dirs := make([]string, 0)
	for _, fi := range fis {
		name := fi.Name()
		if fi.IsDir() {
			dir := filepath.Clean(path + sep + name)
			dirs = append(dirs, dir)
		} else {
			if judgeFunc(pattern, name) {
				return filepath.Clean(path + sep + name), nil
			}
		}
	}

	if len(dirs) != 0 {
		for _, dir := range dirs {
			filename, err := Find(dir, pattern, judgeFunc)
			if len(filename) != 0 && err == nil {
				return filename, nil
			}
		}
	}

	return "", errors.New("file not found")
}
