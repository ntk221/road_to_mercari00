package main

import (
	"path/filepath"
	"os"
)


// func ReadDir(name string) ([]DirEntry, error)
// func Walk(root string, fn WalkFunc) error

// あるディレクトリの配下のファイル一覧を取得する
func listFiles(dirPath string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
