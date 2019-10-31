package utils

import "path"

func GetSuffix(filename string) string {
	onlyFilenameWithSuffix := path.Base(filename)
	suffix := path.Ext(onlyFilenameWithSuffix)
	return suffix
}
