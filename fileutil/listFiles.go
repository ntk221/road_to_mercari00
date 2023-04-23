package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

// func ReadDir(name string) ([]DirEntry, error)
// func Walk(root string, fn WalkFunc) error

// あるディレクトリの配下のファイル一覧を取得する
func ListFiles(dirPath string) ([]string, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("error: %s: no such file or directory", dirPath)
	}
	files := []string{}
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
