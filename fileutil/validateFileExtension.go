package fileutil

import "path/filepath"

func ValidateFileExtension(filePath string, ext []string) bool {
	for _, e := range ext {
		if filepath.Ext(filePath) == e {
			return true
		}
	}
	return false
}
